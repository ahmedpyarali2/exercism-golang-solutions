// Package reverse provides functionality to reverse strings
package reverse

// Reverse reverses a given string input
func Reverse(s string) string {
	// Check if the given input is empty. If it is, we return an empty string
	if s == "" {
		return ""
	}

	// Convert the input string into slice of runes
	sRune := []rune(s)

	// Iterate with two pointers, one from the start and second
	// from the end. Iterate until both pointers collide and swap
	// the values on each run.
	for i, j := 0, len(sRune)-1; i < j; i, j = i+1, j-1 {
		sRune[i], sRune[j] = sRune[j], sRune[i]
	}

	// return the reversed rune slice as a string
	return string(sRune)
}
