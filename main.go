package main

import (
	"github.com/khuang0312/word_frequency/textsource"
	"github.com/khuang0312/word_frequency/utils"
	"io"
	"os"
	"sync"
)

var Links = []string{
	"https://norvig.com/big.txt",
	"https://bereanbible.com/bsb.txt",
	"https://readersbible.com/brb.txt",
	"https://openbible.com/textfiles/asv.txt",
	"https://openbible.com/textfiles/akjv.txt",
	"https://openbible.com/textfiles/cpdv.txt",
	"https://openbible.com/textfiles/dbt.txt",
	"https://openbible.com/textfiles/drb.txt",
	"https://openbible.com/textfiles/erv.txt",
	"https://openbible.com/textfiles/jps.txt",
	"https://openbible.com/textfiles/kjv.txt",
	"https://openbible.com/textfiles/slt.txt",
	"https://openbible.com/textfiles/wbt.txt",
	"https://openbible.com/textfiles/web.txt",
	"https://openbible.com/textfiles/ylt.txt",
}

func CountWordFrequency(r io.Reader) [][]string {
	ts := textsource.NewTextSource(r)
	records := ts.GetSortedRecords()
	return records
}

func processFile(p string) (string, [][]string) {
	f, err := utils.OpenFile(p)
	if err != nil {
		// log error
	}
	defer f.Close()
	tmpFilename := f.Name()
	records := CountWordFrequency(f)
	return tmpFilename, records
}

func IterativeVersion() {
	for _, url := range Links {
		tmpFilename, records := processFile(url)
		utils.WriteToCSV(utils.GetFilename(url), records)
		err := os.Remove(tmpFilename)
		if err != nil {
			// log error
		}
	}

}

func ConcurrentVersion() {
	var wg sync.WaitGroup
	for _, url := range Links {
		wg.Add(1)
		// we need params so goroutine won't end up using same value
		go func(path string) {
			defer wg.Done()
			tmpFilename, records := processFile(url)
			utils.WriteToCSV(utils.GetFilename(url), records)
			err := os.Remove(tmpFilename)
			if err != nil {
				// log error
			}
		}(url)
	}
	wg.Wait()
}

func main() {
	// IterativeVersion()
	ConcurrentVersion()

	utils.CleanupOutputData()

	// download file as temporary file (optional)
	// we download if given an url
	// otherwise check if it exists

	// make textsource from existing file
	// populating the word frequency map
	// get sorted records
	// write it to output

	// deleting all temporary files created

}
