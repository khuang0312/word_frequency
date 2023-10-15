package main

import (
	"testing"
	"os"
)

func setup() {
	// download all the test files
	for url, filename := range Links {
		err := DownloadFile(url, filename)
		if err != nil {
			panic(err)
		}
    }
}

func teardown() {
	// remove all the files
	for _, filename := range Links {
		err := os.Remove(filename)
		if err != nil {
			panic(err)
		}
    }
}



func BenchmarkGetWordFrequencyPerFile(b *testing.B) {
	for _, filename := range Links {
        b.Run(filename, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
				GetWordFrequency(filename)
            }
        })
    }
}

// func BenchmarkGetWordFrequencyForAllFiles(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//         // Parallelized version here
//     }
// }


func TestMain(m *testing.M) {
	setup()
    code := m.Run()
	teardown()
    os.Exit(code)
}