// Package etl provides a functionality to convert legacy scrable score data format
// into a new data format.
package etl

import "strings"

// Transform returns the scores in a new format.
func Transform(scores map[int][]string) map[string]int {
	// initialize a nil map for the transformed scores
	transformedScores := map[string]int{}

	// iterate over each key in the map
	for k, v := range scores {
		// iterate over each element in the value
		for _, s := range v {
			// create an entry in the transformedScores
			transformedScores[strings.ToLower(s)] = k
		}
	}

	// return the transformed scores
	return transformedScores
}
