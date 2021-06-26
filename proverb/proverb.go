// Package proverb provides functionality to generate proverbial rhymes.
package proverb

import "fmt"

// Proverb generates the proverbial rhymes based on the given input.
func Proverb(rhyme []string) []string {
	proverbialRhymes := []string{}

	if len(rhyme) == 0 {
		return proverbialRhymes
	}

	//	1.	We start from index 1 (2nd index) of the given rhyme input and we iterate
	//	  	until the end of the list.
	//	2.	On each iteration we use the current and previous index to generate a rhyme.
	for i := 1; i < len(rhyme); i++ {
		proverbialRhymes = append(proverbialRhymes, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}

	//	3.	At the end we generate the last rhyme using the very first index of the input.
	proverbialRhymes = append(proverbialRhymes, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverbialRhymes
}
