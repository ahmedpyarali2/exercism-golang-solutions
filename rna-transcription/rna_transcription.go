// Package strand contains functionality for conversions between DNA and RNA strands
package strand

var transcribeMap = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA converts the given DNA strand into an RNA strands by replacing each nucleotide in the
// DNA strand with its complement.
func ToRNA(dna string) string {
	rna := []byte{}
	for _, v := range dna {
		rna = append(rna, transcribeMap[byte(v)])
	}

	return string(rna)
}
