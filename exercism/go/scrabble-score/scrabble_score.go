//Package scrabble it's a package to calculate the score of a word
package scrabble

var letters = make(map[string]int)

func init() {
	// A, E, I, O, U, L, N, R, S, T -> 1
	letters["A"] = 1
	letters["E"] = 1
	letters["I"] = 1
	letters["O"] = 1
	letters["U"] = 1
	letters["L"] = 1
	letters["N"] = 1
	letters["R"] = 1
	letters["S"] = 1
	letters["T"] = 1

	// D, G -> 2
	letters["D"] = 2
	letters["G"] = 2

	// B, C, M, P -> 3
	letters["D"] = 3
	letters["C"] = 3
	letters["M"] = 3
	letters["P"] = 3

	// F, H, V, W, Y -> 4
	letters["F"] = 4
	letters["H"] = 4
	letters["V"] = 4
	letters["W"] = 4
	letters["Y"] = 4

	// K -> 5
	letters["K"] = 5

	// J, X -> 8
	letters["J"] = 8
	letters["X"] = 8

	// Q, Z -> 10
	letters["Q"] = 10
	letters["Z"] = 10
}

//Score calculates the score of a word
func Score(word string) (points int) {

	return points
}
