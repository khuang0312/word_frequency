package main

import (
	"testing"
	"gotest.tools/v3/assert"
	"reflect" // needed for deepEqual
)

func TestTokenizeDoesDistinguishCapitalization(t *testing.T) {
	result := Tokenize("the The tHe thE")
	expected := []string{"the", "the", "the", "the"}
	assert.Assert(t, reflect.DeepEqual(result, expected))
}

func TestTokenizeDoesNotParseNumbers(t *testing.T) {
	result := Tokenize("god made 2 great lights")
	expected := []string{"god", "made", "great", "lights"}
	assert.Assert(t, reflect.DeepEqual(result, expected))
}

func TestTokenizeParsesApostrophes(t *testing.T) {
	result := Tokenize("he'd, I'd, I'll, shan't")
	expected := []string{"he'd", "i'd", "i'll", "shan't"}
	assert.Assert(t, reflect.DeepEqual(result, expected))
}

func TestTokenizeParsesHyphens(t *testing.T) {
	result := Tokenize("The state-of-the-art design")
	expected := []string{"the", "state", "of", "the", "art", "design"}
	assert.Assert(t, reflect.DeepEqual(result, expected))
}

func TestTokenizeIgnoresPunctuation(t *testing.T) {
	result := Tokenize("Sam: This is not me,")
	expected := []string{"sam", "this", "is", "not", "me"}
	assert.Assert(t, reflect.DeepEqual(result, expected))
}


