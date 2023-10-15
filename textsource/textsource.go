package textsource

import (
	"bufio"
	// "encoding/csv"
	// "fmt"
	"io"
	// "net/http"
	// "os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// shoved in here to avoid circular dependencies from
// TextSource trying to import Tokenize and test files importing TextSource
func Tokenize(line string) []string {
	// not planning on distinguishing case
	line = strings.ToLower(line)

	// replace special apostrophe with regular one
	line = strings.ReplaceAll(line, "’", "'")

	// -1 indicates return all matches
	return regexp.MustCompile(`[[:alpha:]]+(['])?[[:alpha:]]+`).FindAllString(line, -1)
}

// This has its own package so members aren't visible outside it
type TextSource struct {
	// want this to be generalized beyond files
	// so these members aren't visible outside package
	reader           io.Reader
	wordFrequencyMap map[string]int
}

func NewTextSource(r io.Reader) *TextSource {
	ts := new(TextSource)
	ts.reader = r
	ts.wordFrequencyMap = make(map[string]int)
	ts.populateWordFrequencyMap()
	return ts
}

// should only be called once during construction
func (ts *TextSource) populateWordFrequencyMap() error {
	// prevents segmentation faults from trying to scan from nil
	if ts.reader != nil {
		scanner := bufio.NewScanner(ts.reader)
		for scanner.Scan() {
			line := scanner.Text()
			tokens := Tokenize(line)
			for _, t := range tokens {
				ts.wordFrequencyMap[t] += 1
			}
		}

		// if scanner.Scan() errors, we break out of loop
		if scanner.Err() != nil {
			return scanner.Err()
		}
		// fmt.Println(tf.url, len(tf.wordFrequencyMap))
	}
	return nil
}

func (ts *TextSource) WordFrequencyMap() map[string]int {
	return ts.wordFrequencyMap
}

func (ts *TextSource) GetSortedRecords() [][]string {
	// reflect just becomes too confusing and hard to read
	// maps.keys() is still experimental
	words := []string{}
	for key, _ := range ts.wordFrequencyMap {
		words = append(words, key)
	}
	slices.Sort(words)

	records := [][]string{}
	for _, w := range words {
		frequency := strconv.Itoa(ts.wordFrequencyMap[w])
		records = append(records, []string{w, frequency})
	}

	return records
}
