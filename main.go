package main

import (
	"github.com/khuang0312/word_frequency/utils"
	"sync"
)

var Links = map[string]string{
	"https://norvig.com/big.txt":               "big.txt",
	"https://bereanbible.com/bsb.txt":          "bsb.txt",
	"https://readersbible.com/brb.txt":         "brb.txt",
	"https://openbible.com/textfiles/asv.txt":  "asv.txt",
	"https://openbible.com/textfiles/akjv.txt": "akjv.txt",
	"https://openbible.com/textfiles/cpdv.txt": "cpdv.txt",
	"https://openbible.com/textfiles/dbt.txt":  "dbt.txt",
	"https://openbible.com/textfiles/drb.txt":  "drb.txt",
	"https://openbible.com/textfiles/erv.txt":  "erv.txt",
	"https://openbible.com/textfiles/jps.txt":  "jps.txt",
	"https://openbible.com/textfiles/kjv.txt":  "kjv.txt",
	"https://openbible.com/textfiles/slt.txt":  "slt.txt",
	"https://openbible.com/textfiles/wbt.txt":  "wbt.txt",
	"https://openbible.com/textfiles/web.txt":  "web.txt",
	"https://openbible.com/textfiles/ylt.txt":  "ylt.txt",
}

func IterativeVersion() {
	for url, filename := range Links {
		err := utils.DownloadFile(url, filename)
		if err != nil {
			panic(err)
		}
		frequencyMap, _ := utils.GetWordFrequencyMap(filename)
		records := utils.GetSortedRecordsFromFrequencyMap(frequencyMap)
		utils.WriteToCSV(filename+".csv", records)
		utils.RemoveFile(filename)
	}
}

func ConcurrentVersion() {
	var wg sync.WaitGroup

	for url, filename := range Links {
		wg.Add(1)
		// we need params so goroutine won't end up using same value
		go func(url string, filename string) {
			defer wg.Done()
			utils.DownloadFile(url, filename)
			frequencyMap, _ := utils.GetWordFrequencyMap(filename)
			records := utils.GetSortedRecordsFromFrequencyMap(frequencyMap)
			utils.WriteToCSV(filename+".csv", records)
			utils.RemoveFile(filename)
		}(url, filename)
	}
	wg.Wait()
}

func main() {
	// records := [][]string{
	// 	{"alan", "1"},
	// 	{"james", "12"},
	// }

	// fmt.Println(records)

	// err := utils.WriteToCSV("x.csv", records)

	ConcurrentVersion()
}
