// Package hamming provides useful functionality for calculating the hamming distance for the
// DNA strands.
package hamming

import (
	"errors"
)

// Distance calculates the Hamming Distance between the two given DNA strands.
func Distance(a, b string) (int, error) {
	// we first convert the strings into slice of runes to handle unicode
	aR := []rune(a)
	bR := []rune(b)

	// We return an error in case the two strands of the DNA are not same length.
	if len(aR) != len(bR) {
		return 0, errors.New("cannot calculate distance between strings of different lengths")
	}

	distance := 0

	// We iterate over each DNA codon in the two strands and check if both of them are same.
	// If both the codons are not same, we increment the hamming distance.
	for i, r := range aR {
		if r != bR[i] {
			distance++
		}
	}

	return distance, nil
}
