// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
	"strings"
	"unicode"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
	wordFreq := map[string]int{}
	text = strings.ToLower(text)
	isWord := false
	var word strings.Builder
	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == '\'' {
			if char == '\'' {
				continue
			}
			if !isWord {
				isWord = true
			}
			word.WriteRune(char)
		} else {
			if isWord {
				wordFreq[word.String()]++
			}
			isWord = false
			word = strings.Builder{}
		}
	}
	if isWord {
		wordFreq[word.String()]++
	}
	return wordFreq
} 