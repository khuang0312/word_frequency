package utils

import (
	"regexp"
	"strings"
	"fmt"
)

func Tokenize(line string) []string {
	// not planning on distinguishing case
	line = strings.ToLower(line)

	// replace special apostrophe with regular one
	line = strings.ReplaceAll(line, "’", "'")

	// -1 indicates return all matches
	return regexp.MustCompile(`[[:alpha:]]+(['])?[[:alpha:]]+`).FindAllString(line, -1)
}


func PathToFileName(path string) (string, error) {
	// make sure that we get the last slash or backslash
	lastSlashIndex := -1
	if (strings.LastIndex(path, "/") != -1) {
		lastSlashIndex = strings.LastIndex(path, "/")
	} else if (strings.LastIndex(path, "\\") != -1 ) {
		lastSlashIndex = strings.LastIndex(path, "\\")
	}
	
	if (lastSlashIndex == -1) {
		return "", fmt.Errorf("Path '%s' is not a valid filepath or URL", path)
	}

	filename := path[lastSlashIndex+1:]
	return filename, nil
}

func FilenameToSuffix(filename string) string {
	_, after, _ := strings.Cut(filename, ".")
	return after
}
