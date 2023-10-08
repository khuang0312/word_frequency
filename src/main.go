package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func DownloadFile(url string, filepath string, ) error {
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


func main() {
	links := map[string]string{
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

	for url, fileName := range links {
        fmt.Println(url, fileName)

		err := DownloadFile(url, fileName)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + fileName)
		break;
    }

	

}