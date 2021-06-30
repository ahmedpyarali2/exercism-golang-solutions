// Package isogram provides functionality for checking the isograms
package isogram

import (
	"sort"
	"strings"
)

// IsIsogram checks if the given string is an isogram
// A string is considered to be isogram if there are no repeating
// character in the entire string.
func IsIsogram(s string) bool {
	/*
		We could have chosen to keep a dictionary of characters and track their count,
		but this would result in extra space.
		Another approach would be to sort the entire string and then check, while iterating,
		if any character is repeating.
	*/

	// convert the given string to lower case and then into a slice of runes for handling unicode
	s = strings.ToLower(s)
	sRune := []rune(s)

	// sort the entire slice of rune  - assuming the complexity of sorting is(O(nlogn))
	sort.Slice(sRune, func(i, j int) bool { return sRune[i] < sRune[j] })

	// while iterating over each rune, check if the current rune is same as the
	// previous one, if it is, return false
	for i := 1; i < len(sRune); i++ {
		if sRune[i] != '-' && sRune[i] != ' ' {
			if sRune[i-1] == sRune[i] {
				return false
			}
		}
	}

	// the string is an isogram
	return true
}
