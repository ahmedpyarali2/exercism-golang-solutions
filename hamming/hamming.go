// Package hamming provides useful functionality for calculating the hamming distance for the
// DNA strands.
package hamming

import "errors"

// Distance calculates the Hamming Distance between the two given DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot calculate distance between strings of different lengths")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
