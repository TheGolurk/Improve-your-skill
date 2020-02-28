// Package diffsquares calculates the differences between SquareOfSum and SumOfSquares
package diffsquares

// SquareOfSum sum some natural numbers
func SquareOfSum(number int) (result int) {

	for i := 1; i <= number; i++ {
		result += i
	}
	result *= result

	return result
}

// SumOfSquares sum of square of some numbers
func SumOfSquares(number int) (result int) {

	for i := 1; i <= number; i++ {
		result += i * i
	}

	return result
}

// Difference between the square of the sum and the sum of the squares
func Difference(number int) int {
	var resultSquareOfSum, resultSumOfSquare int

	for i := 1; i <= number; i++ {
		resultSquareOfSum += i
		resultSumOfSquare += i * i
	}
	resultSquareOfSum *= resultSquareOfSum

	return resultSquareOfSum - resultSumOfSquare
}
