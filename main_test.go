package main

import (
	"github.com/khuang0312/word_frequency/utils"
	"os"
	"testing"
)

func teardown() {
	utils.CleanupOutputData()
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
	code := m.Run()
	teardown()
	os.Exit(code)
}
