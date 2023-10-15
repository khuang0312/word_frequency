package main

import (
	"os"
	"testing"
)

func teardown() {
	// remove all the files
	err := os.RemoveAll("output")
	if err != nil {
		panic(err)
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
	code := m.Run()
	teardown()
	os.Exit(code)
}
