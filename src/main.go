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


func GetStringsFromLine(line string) []string {
	// -1 indicates return all matches
	// words only have Latin letters and apostrophes

	// write unit tests for this
	// line := "Genesis 1:16	And God made two great lights: a greater light to rule the day; and a lesser light to rule the night: and the stars."
	// line2 := "And Abram and Nachor married wives: the name of Abram’s wife was Sarai: and the name of Nachor’s wife, Melcha, the daughter of Aran, father of Melcha, and father of Jescha."

	return regexp.MustCompile(`[[:alpha:]'’]+`).FindAllString(line, -1)
}

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

func ReadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
        return err
    }
	defer file.Close();

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
		lineStrings := strings.Fields(line)
		fmt.Println(lineStrings)
		
    }
    return scanner.Err()

}


func main() {
	// links := map[string]string{
	// 	"https://norvig.com/big.txt": "big.txt",
	// 	"https://bereanbible.com/bsb.txt": "bsb.txt",
	// 	"https://readersbible.com/brb.txt": "brb.txt",
	// 	"https://openbible.com/textfiles/asv.txt": "asv.txt",
	// 	"https://openbible.com/textfiles/akjv.txt" : "akjv.txt",
	// 	"https://openbible.com/textfiles/cpdv.txt": "cpdv.txt",
	// 	"https://openbible.com/textfiles/dbt.txt": "dbt.txt",
	// 	"https://openbible.com/textfiles/drb.txt": "drb.txt",
	// 	"https://openbible.com/textfiles/erv.txt": "erv.txt",
	// 	"https://openbible.com/textfiles/jps.txt": "jps.txt",
	// 	"https://openbible.com/textfiles/kjv.txt": "kjv.txt",
	// 	"https://openbible.com/textfiles/slt.txt": "slt.txt",
	// 	"https://openbible.com/textfiles/wbt.txt": "wbt.txt",
	// 	"https://openbible.com/textfiles/web.txt": "web.txt",
	// 	"https://openbible.com/textfiles/ylt.txt": "ylt.txt",
	// }

	// for url, fileName := range links {
    //     fmt.Println(url, fileName)

	// 	err := DownloadFile(url, fileName)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Downloaded: " + fileName)
	// 	ReadFile(fileName)
	// 	break;
    // }
}