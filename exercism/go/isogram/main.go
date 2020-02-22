package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsIsogram("isogram"))
}

// IsIsogram leadasdasdsdasds
func IsIsogram(word string) (is bool) {
	for _, letter := range word {

		if con := strings.Count(word, string(letter)); con > 1 {
			return false
		}

	}

	return true
}
