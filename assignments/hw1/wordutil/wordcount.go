package wordutil

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {

	m := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		word = strings.ToLower(word)
		_, exists := m[word]
		if exists {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	return m
}
