package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func DownloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	fmt.Println("Downloaded: " + filepath)

	return err
}

func RemoveFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	fmt.Println("Removed " + filename)
	return nil
}

func GetWordFrequencyMap(filepath string) (map[string]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	frequencyMap := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := Tokenize(line)
		for _, t := range tokens {
			frequencyMap[t] += 1
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	fmt.Println(filepath, len(frequencyMap))
	return frequencyMap, nil
}

func GetSortedRecordsFromFrequencyMap(frequencyMap map[string]int) [][]string {
	// no good way to iterate through maps in a sorted order as of this time
	words := reflect.ValueOf(frequencyMap).MapKeys()
	slices.SortFunc(words, func(a, b reflect.Value) int {
		return strings.Compare(a.String(), b.String())
	})

	records := [][]string{}
	for _, w := range words {
		word := w.String()

		records = append(records, []string{word, strconv.Itoa(frequencyMap[word])})
	}

	return records
}

func WriteToCSV(filename string, records [][]string) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println(records)

	w := csv.NewWriter(f)
	defer w.Flush()

	headers := []string{"word", "count"}
	w.Write(headers)

	for _, record := range records {
		w.Write(record)
	}
	
	return err
}
