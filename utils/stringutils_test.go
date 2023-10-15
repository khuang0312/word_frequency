package utils

import (
	"gotest.tools/v3/assert"
	"testing"
	"fmt"
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

func TestPathToFileNameWithURL(t *testing.T) {
	url := "https://norvig.com/big.txt"
	result, _ := PathToFileName(url)
	assert.Equal(t, result, "big.txt")
}

func TestPathToFileNameWithRootURL(t *testing.T) {
	url := "https://norvig.com/"
	result, _ := PathToFileName(url)
	assert.Equal(t, result, "")
}

func TestPathToFileNameWithLinuxFilePath(t *testing.T) {
	path := "/big.txt"
	result, _ := PathToFileName(path)
	assert.Equal(t, result, "big.txt")
}

func TestPathToFileNameWithWindowsFilePath(t *testing.T) {
	path := "C:Projects\\apilibrary\\apilibrary.sln"
	result, _ := PathToFileName(path)
	assert.Equal(t, result, "apilibrary.sln")
}

func TestPathToFileNameWithoutSlashOrBackSlash(t *testing.T) {
	path := "fail"
	_, err := PathToFileName(path)
	assert.Error(t, err, fmt.Sprintf("Path '%s' is not a valid filepath or URL", path))
}

func TestFileNameToSuffixWithOneDot(t *testing.T) {
	filename := "x.csv"
	assert.Equal(t, FilenameToSuffix(filename), "csv")
}

func TestFileNameToSuffixWithTwoDots(t *testing.T) {
	filename := "x.tar.gz"
	assert.Equal(t, FilenameToSuffix(filename), "tar.gz")
}

func TestFileNameToSuffixWithoutPeriod(t *testing.T) {
	filename := "xcsv"
	assert.Equal(t, FilenameToSuffix(filename), "")
}
