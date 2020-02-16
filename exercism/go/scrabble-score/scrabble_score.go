//Package scrabble it's a package to calculate the score of a word
package scrabble

import "unicode"

var letters = map[rune]int{
	// A, E, I, O, U, L, N, R, S, T -> 1
	'A': 1,
	'E': 1,
	'I': 1,
	'O': 1,
	'U': 1,
	'L': 1,
	'N': 1,
	'R': 1,
	'S': 1,
	'T': 1,

	// D, G -> 2
	'D': 2,
	'G': 2,

	// B, C, M, P -> 3
	'B': 3,
	'C': 3,
	'M': 3,
	'P': 3,

	// F, H, V, W, Y -> 4
	'F': 4,
	'H': 4,
	'V': 4,
	'W': 4,
	'Y': 4,

	// K -> 5
	'K': 5,

	// J, X -> 8
	'J': 8,
	'X': 8,

	// Q, Z -> 10
	'Z': 10,
	'Q': 10,
}

//Score calculates the score of a word
func Score(word string) (points int) {

	for _, value := range word {
		points += letters[unicode.ToUpper(value)]
	}

	return points
}
