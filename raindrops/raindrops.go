// Package raindrops provides functionality for conversion of numbers to raindrop sounds.
package raindrops

import (
	"strconv"
)

// Convert converts a given number into string that contains raindrop sounds
// corresponding to certain potential factors.
func Convert(number int) string {
	raindropSound := ""

	// add a "Pling" sound if the number is a factor of 3
	if number%3 == 0 {
		raindropSound += "Pling"
	}

	// add a "Plang" sound if the number is a factor of 5
	if number%5 == 0 {
		raindropSound += "Plang"
	}

	// add a "Plong" sound if the number is a factor of 7
	if number%7 == 0 {
		raindropSound += "Plong"
	}

	// add the number itself if the number is not a factor of either 3, 5 or 7
	if raindropSound == "" {
		return strconv.Itoa(number)
	}

	return raindropSound
}
