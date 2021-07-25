// Package anagram provides functionality for identifying
// anagrams.
package anagram

import (
	"sort"
	"strings"
)

type runeSlice []rune

func (r runeSlice) Len () int { return len(r) }
func (r runeSlice) Less (i, j int) bool { return r[i] < r[j] }
func (r runeSlice) Swap (i, j int) { r[i], r[j] = r[j], r[i] }

// sortRune sorts the given rune slice
func sortRune(entry runeSlice) []rune {
	localEntry := make(runeSlice, len(entry), len(entry))
	copy(localEntry, entry)

	sort.Sort(localEntry)
	return localEntry
}

// isRuneSame compares the two rune slices and checks if the
// two rune slices are same or not
func isRuneSame(a, b runeSlice) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Detect detects if any of the given candidates is an anagram.
// Returns a list of matching candidates.
func Detect(subject string, candidates []string) []string {
	// initialize an empty list of string
	var matchingCandidates []string

	// convert the subject into a slice of rune and sort the slice
	subjectRune := runeSlice(strings.ToLower(subject))
	sortedSubjectRune := sortRune(subjectRune)

	// while iterating over each candidate
	for _, candidate := range candidates {
		// star off with assuming that the candidate is an anagram
		isAnagram := true

		// convert the candidate into a slice of rune and sort it
		candidateRune := runeSlice(strings.ToLower(candidate))
		sortedCandidateRune := sortRune(candidateRune)

		// while iterating over both sorted subject and sorted candidate
		if len(subjectRune) == len(candidateRune) && !isRuneSame(subjectRune, candidateRune) {
			for i := range sortedCandidateRune {
				if sortedSubjectRune[i] != sortedCandidateRune[i] {
					isAnagram = false
					break
				}
			}
		} else {
			isAnagram = false
		}

		if isAnagram {
			matchingCandidates = append(matchingCandidates, candidate)
		}
	}
	// return the matching candidates
	return matchingCandidates
}
