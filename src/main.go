package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
	"bufio"
	"strings"
	"regexp"
)

var Links = map[string]string{
	"https://norvig.com/big.txt": "big.txt",
	"https://bereanbible.com/bsb.txt": "bsb.txt",
	"https://readersbible.com/brb.txt": "brb.txt",
	"https://openbible.com/textfiles/asv.txt": "asv.txt",
	"https://openbible.com/textfiles/akjv.txt" : "akjv.txt",
	"https://openbible.com/textfiles/cpdv.txt": "cpdv.txt",
	"https://openbible.com/textfiles/dbt.txt": "dbt.txt",
	"https://openbible.com/textfiles/drb.txt": "drb.txt",
	"https://openbible.com/textfiles/erv.txt": "erv.txt",
	"https://openbible.com/textfiles/jps.txt": "jps.txt",
	"https://openbible.com/textfiles/kjv.txt": "kjv.txt",
	"https://openbible.com/textfiles/slt.txt": "slt.txt",
	"https://openbible.com/textfiles/wbt.txt": "wbt.txt",
	"https://openbible.com/textfiles/web.txt": "web.txt",
	"https://openbible.com/textfiles/ylt.txt": "ylt.txt",
}

func Tokenize(line string) []string {
	// not planning on distinguishing case
	line = strings.ToLower(line)

	// replace special apostrophe with regular one
	line = strings.ReplaceAll(line, "’", "'")

	// -1 indicates return all matches
	return regexp.MustCompile(`[[:alpha:]]+(['])?[[:alpha:]]+`).FindAllString(line, -1)
}

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
	return err
}

func GetWordFrequency(filepath string) (error, map[string]int) {
	file, err := os.Open(filepath)
	if err != nil {
        return err, nil
    }
	defer file.Close();

	frequencyDict := map[string]int{}

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		tokens:= Tokenize(line)
		for _, t := range tokens {
			frequencyDict[t] += 1
		}	
    }
	if scanner.Err() != nil {
		return scanner.Err(), nil
	}

	return nil, frequencyDict
}


func main() {
	

	for url, filename := range Links {
        fmt.Println(url, filename)

		err := DownloadFile(url, filename)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + filename)
		fmt.Println(GetWordFrequency(filename))
		break;
		
    }
}