// Package accumulate provides functionality related to collection accumulation
package accumulate

// Accumulate provides a functionality to apply the given accumulator function on
// the given array of string.
func Accumulate(s []string, accumulator func(string) string) []string {
	// make a new array to return
	modified := make([]string, len(s))

	// iterate over all the elements of the given array of string
	for i := range s {
		// on each element apply the accumulator function and add it to the array to return
		modified[i] = accumulator(s[i])
	}

	return modified
}
