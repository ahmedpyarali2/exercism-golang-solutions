// Package wordcount provides functionality for word frequency calculation
package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile(`(\S+\'\S+)|([0-9a-zA-Z]+)`)

func WordCount(sentence string) Frequency {
	// create an empty Frequency store
	frequencyStore := Frequency{}

	// create a replace that will replace escape characters i-e: "\t", "\n", "\r"
	replacer := strings.NewReplacer("\n", " ", "\t", " ", "\r", " ")
	sentence = replacer.Replace(sentence)

	// perform a regular expression search and create tokens for matching patterns
	tokens := re.FindAllString(sentence, -1)

	// while iterating over the list of tokens
	for _, t := range tokens {
		// convert the token to lowercase
		t = strings.ToLower(t)

		// if the token does not exist in the frequency store, initialize the score from 1
		if _, ok := frequencyStore[t]; ok {
			frequencyStore[t]++
		} else {
			// else increment the count
			frequencyStore[t] = 1
		}
	}

	// return the frequency store
	return frequencyStore
}
