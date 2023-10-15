package main

import (
	"github.com/khuang0312/word_frequency/utils"
	"os"
	"testing"
)

func setup() {
	// download all the test files
	for url, filename := range Links {
		err := utils.DownloadFile(url, filename)
		if err != nil {
			panic(err)
		}
	}
}

func teardown() {
	// remove all the files
	for _, filename := range Links {
		err := utils.RemoveFile(filename)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkIterativeVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativeVersion()
	}
}

func BenchmarkConcurrentVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentVersion()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
