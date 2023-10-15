package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"
)

type TextFile struct {
	url              string
	tmpFilename      string
	wordFrequencyMap map[string]int
}

func NewTextFile(url string) *TextFile {
	tf := new(TextFile)
	tf.url = url
	tf.wordFrequencyMap = make(map[string]int)
	return tf
}

// we want to read from this file object after...
func (tf *TextFile) DownloadFile() error {
	resp, err := http.Get(tf.url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	// because we'll just be deleting this soon anyways
	out, err := os.CreateTemp("", "")
	defer out.Close()
	if err != nil {
		return err
	}

	// save file name to access later
	tf.tmpFilename = out.Name()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (tf *TextFile) RemoveFile() error {
	err := os.Remove(tf.tmpFilename)
	if err != nil {
		return err
	}

	// fmt.Println("Removed " + tf.tmpFilename)
	return nil
}

func (tf *TextFile) PopulateWordFrequencyMap() error {
	reader, err := os.Open(tf.tmpFilename)
	defer reader.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := Tokenize(line)
		for _, t := range tokens {
			tf.wordFrequencyMap[t] += 1
		}
	}

	// if scanner.Scan() errors, we break out of loop
	if scanner.Err() != nil {
		return scanner.Err()
	}
	// fmt.Println(tf.url, len(tf.wordFrequencyMap))
	return nil
}

func (tf *TextFile) GetSortedRecords() [][]string {
	// reflect just becomes too confusing and hard to read
	// didn't want to write a loop just to iterate through a map
	words := []string{}
	for key, _ := range tf.wordFrequencyMap {
		words = append(words, key)
	}
	slices.Sort(words)

	records := [][]string{}
	for _, w := range words {
		records = append(records, []string{w, strconv.Itoa(tf.wordFrequencyMap[w])})
	}

	return records
}

func (tf *TextFile) GetCSVFilename() (string, error) {
	filename, err := PathToFilename(tf.url)
	if err != nil {
		return "", err
	}

	filename = FilenameToPrefix(filename)
	filename = fmt.Sprintf("%s.csv", filename)
	return filename, nil
}

// this isn't a method because this is generalizable beyond TextFiles
func WriteToCSV(filename string, records [][]string) error {
	// 660 is rw for user and groups
	err := os.Mkdir("output", 660)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// truncates existing files by default
	f, err := os.Create("output/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	headers := []string{"word", "count"}
	w.Write(headers)

	for _, record := range records {
		w.Write(record)
	}

	return err
}
