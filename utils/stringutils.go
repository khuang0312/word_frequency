package utils

import (
	// "fmt"
	// "strings"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

// so we don't have to store a mapping of URLs to their file names
func FilepathToFilename(path string) string {
	// make sure that we get the last slash or backslash
	_, file := filepath.Split(path)
	return file
}

func UrlToFilename(url string) string {
	// make sure that we get the last slash or backslash
	_, file := path.Split(url)
	return file
}

func IsValidURL(p string) bool {
	// this is what we want...
	// [scheme:][//[userinfo@]host][/]path
	u, err := url.Parse(p)
	return (err == nil &&
		u.Scheme != "" &&
		u.Host != "" &&
		path.Ext(u.Path) == ".txt" &&
		u.RawQuery == "" &&
		u.RawFragment == "")
}

func IsValidFile(path string) bool {
	if filepath.Ext(path) != ".txt" {
		return false
	}

	file, err := os.Open(path)
	if err != nil {
		return false
	}

	defer file.Close()

	return true
}
