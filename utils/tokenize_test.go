package utils

import (
	"github.com/khuang0312/word_frequency/textsource"
	"gotest.tools/v3/assert"
	"testing"
)

func TestTokenizeDoesDistinguishCapitalization(t *testing.T) {
	result := textsource.Tokenize("the The tHe thE")
	expected := []string{"the", "the", "the", "the"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeDoesNotParseNumbers(t *testing.T) {
	result := textsource.Tokenize("god made 2 great lights")
	expected := []string{"god", "made", "great", "lights"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeParsesApostrophes(t *testing.T) {
	result := textsource.Tokenize("he'd, I'd, I'll, shan't")
	expected := []string{"he'd", "i'd", "i'll", "shan't"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeParsesHyphens(t *testing.T) {
	result := textsource.Tokenize("The state-of-the-art design")
	expected := []string{"the", "state", "of", "the", "art", "design"}
	assert.DeepEqual(t, result, expected)
}

func TestTokenizeIgnoresPunctuation(t *testing.T) {
	result := textsource.Tokenize("Sam: This is not me,")
	expected := []string{"sam", "this", "is", "not", "me"}
	assert.DeepEqual(t, result, expected)
}
