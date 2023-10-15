package main

import (
	"github.com/khuang0312/word_frequency/utils"
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

func CountWordFrequency(url string) {
	tf := utils.NewTextFile(url)
	err := tf.DownloadFile()
	if err != nil {
		panic(err)
	}

	err = tf.PopulateWordFrequencyMap()
	if err != nil {
		panic(err)
	}

	// we don't need the file anymore
	err = tf.RemoveFile()
	if err != nil {
		panic(err)
	}

	records := tf.GetSortedRecords()
	filename, err := tf.GetCSVFilename()
	if err != nil {
		panic(err)
	}

	utils.WriteToCSV(filename, records)
}

func IterativeVersion() {
	for _, url := range Links {
		CountWordFrequency(url)
	}
}

func ConcurrentVersion() {
	var wg sync.WaitGroup
	for _, url := range Links {
		wg.Add(1)
		// we need params so goroutine won't end up using same value
		go func(url string) {
			defer wg.Done()
			CountWordFrequency(url)
		}(url)
	}
	wg.Wait()
}

func main() {
	IterativeVersion()
	ConcurrentVersion()
}
