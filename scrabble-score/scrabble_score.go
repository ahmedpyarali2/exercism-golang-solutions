// Package scrabble provides functionality to calculate scrabble scores on the strings.
package scrabble

import "unicode"

// Score calculates the scrabble score on the given string.
func Score(s string) int {
	var score int

	for _, v := range s {
		switch unicode.ToLower(v) {
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		default:
			score++
		}
	}

	return score
}
