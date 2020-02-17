package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(IsIsogram("isogram"))
}

// IsIsogram leadasdasdsdasds
func IsIsogram(word string) bool {
	for _, letter := range word {
		fmt.Println(letter)
		for index, subLetter := range word {
			fmt.Println(index, subLetter)
			if index > 0 {
				fmt.Println("Dentro de el if: ", index, subLetter)
				if unicode.ToUpper(letter) == unicode.ToUpper(subLetter) {
					fmt.Println("LETRA IGUAL")
					return false
				}
			}
		}
	}

	return true
}
