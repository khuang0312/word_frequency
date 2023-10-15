package utils

import (
	"github.com/khuang0312/word_frequency/textsource"
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

// this test file is not in textsource
// so we avoid accessing private members
// since Golang privacy only takes effect outside package

func TestNewTextSourceCanBeCreatedWithoutReader(t *testing.T) {
	ts := textsource.NewTextSource(nil)
	assert.Equal(t, len(ts.WordFrequencyMap()), 0)
}

func TestTextSourcePopulatesAutomaticallyOnCreation(t *testing.T) {
	reader := strings.NewReader(
		"The quick brown fox jumped. The slow brown fox didn't.")
	ts := textsource.NewTextSource(reader)
	assert.Equal(t, len(ts.WordFrequencyMap()), 7)
}

func TestTextSourceGetSortedRecordsWhenThereIsNoFile(t *testing.T) {
	reader := strings.NewReader(
		"The quick brown fox jumped. The slow brown fox didn't.")
	ts := textsource.NewTextSource(reader)
	expected := [][]string{
		{"brown", "2"},
		{"didn't", "1"},
		{"fox", "2"},
		{"jumped", "1"},
		{"quick", "1"},
		{"slow", "1"},
		{"the", "2"},
	}

	assert.DeepEqual(t, ts.GetSortedRecords(), expected)

}

// func TestTextFileGetCSVFilename(t *testing.T) {

// }

// // records := [][]string{
// // 	{"alan", "1"},
// // 	{"james", "12"},
// // }

// // fmt.Println(records)

// // err := utils.WriteToCSV("x.csv", records)
