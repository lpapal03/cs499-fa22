package wordutil

import (
	"strings"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {

	m := make(map[string][]int)
	var starting_indices []int
	words := strings.Fields(s)

	starting_indices = append(starting_indices, 0)
	for pos, char := range s {
		if char == ' ' {
			starting_indices = append(starting_indices, pos+1)
		}
	}

	for i, word := range words {
		word = strings.ToLower(word)
		m[word] = append(m[word], starting_indices[i])
	}

	return m
}
