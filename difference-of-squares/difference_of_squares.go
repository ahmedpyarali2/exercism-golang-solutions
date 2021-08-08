// Package diffsquares provides functionality to calculate:
//	1. Square of a sum of the first N natural numbers
//	2. Sum of squares of the first N natural numbers
//	3. Difference between 1 and 2
package diffsquares

// SquareOfSum returns the square of the sum of the
// natural numbers from 1..N
func SquareOfSum(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns the sum of all the squares of
// the natural numbers from 1..N
func SumOfSquares(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// Difference returns the difference between square of the sum
// and the sum of the squares for the natural numbers from 1..N
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
