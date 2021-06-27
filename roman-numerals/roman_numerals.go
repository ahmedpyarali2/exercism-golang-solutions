// Package romannumerals provides functionality to convert arabic numbers to Roman Numerals.
// This package implements the solution based on this video: https://www.youtube.com/watch?v=aAYERFf-sMA
package romannumerals

import (
	"errors"
	"strings"
)

var (
	// For optimized solution
	arabicNumerals = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanEquivlent = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
)

// closestMinimum finds the closest minimum value in the list of arabic numerals to the
// given number and returns its index.
// This method performs a binary search and keeps on looking for the closest minimum
// by dividing the search space into half on each iteration.
func closestMinimum(number int) int {
	x, y, mid := 0, len(arabicNumerals), 0

	for {
		newMid := ((x + y) - 1) / 2
		if newMid == mid {
			return mid
		}

		mid = newMid
		if arabicNumerals[mid] == number {
			return mid
		}

		if arabicNumerals[mid] < number {
			x = mid + 1
		} else {
			y = mid
		}
	}

}

// ToRomanNumeral implements a convertor to convert any arabic number (<= 3000) into
// its Roman Numeral form. The solution implemented here is based on this video:
// https://www.youtube.com/watch?v=aAYERFf-sMA
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic <= 0 {
		return "", errors.New("arabic number greater than 3000 and less than 0 not supported")
	}

	var roman strings.Builder

	// iterate until the number is 0.
	for arabic > 0 {
		// find the index of the closest minimum number from the arabic numerals.
		i := closestMinimum(arabic)

		// append its roman equivlent.
		roman.WriteString(romanEquivlent[i])

		// subtract the closest minimum from the actual number.
		arabic -= arabicNumerals[i]
	}

	return roman.String(), nil
}
