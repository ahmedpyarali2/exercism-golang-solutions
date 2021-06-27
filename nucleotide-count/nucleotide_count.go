// Package dna provides functionality related to DNA nucleotides
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides.
type DNA []byte

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	// initialize a new histogram (counter)
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	// we iterate over each nucleotide in DNA
	for _, v := range d {
		// check if its a valid nucleotide, if not, return error
		if _, ok := h[v]; !ok {
			return h, errors.New("invalid nucleotide found")
		}

		// increment the count of the valid nucleotide
		h[v]++

	}

	// return values without any error
	return h, nil
}
