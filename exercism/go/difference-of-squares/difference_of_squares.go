// Package diffsquares calculates the differences between SquareOfSum and SumOfSquares
package diffsquares

// SquareOfSum sum some natural numbers
func SquareOfSum(number int) (result int) {

	result = (number * (number + 1)) / 2

	return result * result
}

// SumOfSquares sum of square of some numbers
func SumOfSquares(number int) (result int) {

	result = (number * (number + 1) * ((number * 2) + 1)) / 6

	return result
}

// Difference between the square of the sum and the sum of the squares
func Difference(number int) int {

	return SquareOfSum(number) - SumOfSquares(number)
}
