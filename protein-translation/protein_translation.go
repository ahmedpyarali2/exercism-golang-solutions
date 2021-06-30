// Package protein provides functionality for RNA translation
package protein

import "errors"

var (
	// ErrStop represents STOP codon
	ErrStop error = errors.New("stop")

	// ErrInvalidBase represents an invalid base exception
	ErrInvalidBase error = errors.New("invalid base")
)

// FromRNA translates the given RNA sequence into proteins.
// The RNA sequence can be broken down into codons and then translated to
// a polypeptide (sequence of proteins).
func FromRNA(rna string) ([]string, error) {
	// we first convert the given string into runes to avoid
	rnaRune := []rune(rna)

	// we check if the length of the given rna sequence is < 3
	if len(rnaRune) < 3 {
		return []string{}, ErrInvalidBase
	}
	// initialize a string slice to return
	polypeptide := make([]string, 0, len(rnaRune)/3)

	for i := 0; i < len(rnaRune)-2; i += 3 {
		// translate the codon into the amino acid
		aminoAcid, err := FromCodon(rna[i : i+3])
		if err != nil {
			// we do not return an error if we receive a STOP codon
			if err == ErrStop {
				err = nil
			}

			// we return the polypeptide and return the error if there is any
			return polypeptide, err
		}

		polypeptide = append(polypeptide, aminoAcid)

	}
	return polypeptide, nil
}

// FromCodon translates the given codon from a RNA sequence,
// into a amino acid.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
