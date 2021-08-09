// Package luhn provides functionality to validate the
// given string based on the luhn algorithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the given string is a valid
// based on the rules of luhn algorithm
//
// Benchmark:
// 	cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// 	BenchmarkValid-12        2647303               421.3 ns/op            80 B/op         17 allocs/op
func Valid(s string) bool {
	// remove spaces from the given string
	s = strings.Replace(s, " ", "", -1)
	// calculate the length of s and sum to verify the string
	lenS, sum := len(s), 0

	// check if the length of the string is less than or equal to 1
	if lenS <= 1 {
		// if it is, return false
		return false
	}

	// tracker for keeping track of every second digit
	isSecondDigit := false

	// start iterating from the right hand side of the given string to the left hand side
	for i := lenS - 1; i >= 0; i-- {
		// we check if the given character is a valid integer and convert it to an integer
		if '0' <= s[i] && s[i] <= '9' {
			currentDigit, _ := strconv.Atoi(string(s[i]))

			// if the current index is divisible by 2 (every second index), we double the value on this index
			if currentDigit > 0 && isSecondDigit {
				currentDigit *= 2

				// if the doubled value is greater than 9, we subtract 9 from it
				if currentDigit > 9 {
					currentDigit -= 9
				}

				isSecondDigit = false
			} else {
				isSecondDigit = true
			}

			// we add the current value to the sum
			sum += currentDigit
		} else {
			// if its not a valid digit we immediately return
			return false
		}

	}

	// if sum is not divisible my 10, its not a valid luhn string
	if sum%10 != 0 {
		return false
	}

	// its a valid luhn string
	return true
}
