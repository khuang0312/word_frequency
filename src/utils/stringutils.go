package utils

import (
	"regexp"
	"strings"
)

func Tokenize(line string) []string {
	// not planning on distinguishing case
	line = strings.ToLower(line)

	// replace special apostrophe with regular one
	line = strings.ReplaceAll(line, "’", "'")

	// -1 indicates return all matches
	return regexp.MustCompile(`[[:alpha:]]+(['])?[[:alpha:]]+`).FindAllString(line, -1)
}
