// Package raindrops is a package to know the factor of a number
package raindrops

import "strconv"

// Convert calculates the factor of a number
func Convert(number int) (result string) {
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		result += strconv.Itoa(number)
	}

	return result
}
