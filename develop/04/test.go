package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDeleteRepeated(t *testing.T) {
	input := []string{"aa", "a", "ab", "ab", "a", "c", "aa"}
	expected := []string{"aa", "a", "ab", "c"}

	result := deleteRepeated(input)
	assert.True(t, reflect.DeepEqual(result, expected))
}

func TestMakeAnagrammDict(t *testing.T) {
	input := []string{"орел", "Катер", "Актер", "рысь", "сырь", "катер", "Сырь", "рысь", "катер", "терка"}
	expected := map[string][]string{
		"катер": {"актер", "катер", "терка"},
		"рысь":  {"рысь", "сырь"},
	}

	result := AnagramDict(input)
	assert.True(t, reflect.DeepEqual(expected, result))
}
