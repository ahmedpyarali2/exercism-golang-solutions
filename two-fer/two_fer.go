// Package twofer implements a single function that generates a twofer string.
package twofer

import "fmt"

// ShareWith generates a string with the given name.
func ShareWith(name string) string {
	// replace the name with `you` if the provided name is empty string.
	if name == "" {
		name = "you"
	}

	// format the final string with the give name.
	// NOTE: could have simply concatenated the strings with `+` but
	// upon running the benchmark test, concatenation approach seemed to be very slow
	// as compared to this string formatting approach. I am not sure if there is any other
	// way that is even faster than this one.
	return fmt.Sprintf("One for %s, one for me.", name)
}
