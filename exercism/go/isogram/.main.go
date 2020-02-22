package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsIsogram("Alphabet"))
}

// IsIsogram leadasdasdsdasds
func IsIsogram(word string) (is bool) {
	wordLow := strings.ToLower(word)
	for _, letter := range wordLow {

		if con := strings.Count(wordLow, string(letter)); con > 1 {
			return false
		}

	}

	return true
}
