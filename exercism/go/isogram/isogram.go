// Package isogram determinates if a word is an isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram is a function to know if a word is an isogram
func IsIsogram(word string) (is bool) {
	wordLow := strings.ToLower(word)
	for _, letter := range wordLow {

		count := strings.Count(wordLow, string(letter))
		if count > 1 && unicode.IsLetter(letter) == true {
			return false
		}

	}

	return true
}
