package utils

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestTokenizeDoesDistinguishCapitalization(t *testing.T) {
	result := Tokenize("the The tHe thE")
	expected := []string{"the", "the", "the", "the"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeDoesNotParseNumbers(t *testing.T) {
	result := Tokenize("god made 2 great lights")
	expected := []string{"god", "made", "great", "lights"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeParsesApostrophes(t *testing.T) {
	result := Tokenize("he'd, I'd, I'll, shan't")
	expected := []string{"he'd", "i'd", "i'll", "shan't"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeParsesHyphens(t *testing.T) {
	result := Tokenize("The state-of-the-art design")
	expected := []string{"the", "state", "of", "the", "art", "design"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeIgnoresPunctuation(t *testing.T) {
	result := Tokenize("Sam: This is not me,")
	expected := []string{"sam", "this", "is", "not", "me"}
	assert.DeepEqual(t, result, expected)
}

func TestPathToFilenameWithURL(t *testing.T) {
	url := "https://norvig.com/big.txt"
	result, _ := PathToFilename(url)
	assert.Equal(t, result, "big.txt")
}

func TestPathToFilenameWithRootURL(t *testing.T) {
	url := "https://norvig.com/"
	result, _ := PathToFilename(url)
	assert.Equal(t, result, "")
}

func TestPathToFilenameWithLinuxFilePath(t *testing.T) {
	path := "/big.txt"
	result, _ := PathToFilename(path)
	assert.Equal(t, result, "big.txt")
}

func TestPathToFilenameWithWindowsFilePath(t *testing.T) {
	path := "C:Projects\\apilibrary\\apilibrary.sln"
	result, _ := PathToFilename(path)
	assert.Equal(t, result, "apilibrary.sln")
}

func TestPathToFilenameWithoutSlashOrBackSlash(t *testing.T) {
	path := "fail"
	_, err := PathToFilename(path)
	assert.Error(t, err, fmt.Sprintf("Path '%s' is not a valid filepath or URL", path))
}

func TestFilenameToSuffixWithOneDot(t *testing.T) {
	filename := "x.csv"
	assert.Equal(t, FilenameToSuffix(filename), "csv")
}

func TestFilenameToSuffixWithTwoDots(t *testing.T) {
	filename := "x.tar.gz"
	assert.Equal(t, FilenameToSuffix(filename), "tar.gz")
}

func TestFilenameToSuffixWithoutPeriod(t *testing.T) {
	filename := "xcsv"
	assert.Equal(t, FilenameToSuffix(filename), "")
}

func TestFilenameToPrefixWithOneDot(t *testing.T) {
	filename := "x.csv"
	assert.Equal(t, FilenameToPrefix(filename), "x")
}

func TestFilenameToPrefixWithTwoDots(t *testing.T) {
	filename := "x.tar.gz"
	assert.Equal(t, FilenameToPrefix(filename), "x")
}

func TestFilenameToPrefixWithoutPeriod(t *testing.T) {
	filename := "xcsv"
	assert.Equal(t, FilenameToPrefix(filename), "xcsv")
}
