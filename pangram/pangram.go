// Package pangram provides functionality for pangram identification
package pangram

import (
	"fmt"
	"github.com/golang-collections/go-datastructures/bitarray"
)

// IsPangram checks if the given sentence in a string is a valid
// pangram or not. A pangram (Greek: παν γράμμα, pan gramma, "every letter") is a
// sentence using every letter of the alphabet at least once. The best known English pangram is:
//		"The quick brown fox jumps over the lazy dog."
func IsPangram(s string) bool {
	/* We simply convert the given string into a byte array and iterate over the byte array.
	Before conversion, we need an array of bits to of size 26 (representing 26 alphabets) to all
	be initialized with a 0 bit. While we iterate, we check if the given character is an alphabet.
	If it is, we mark the bit representing that alphabet in the array as on (1).
	At the end, all of the bits of the bit array must be 1 (on) in order for a string to be pangram.
	*/

	// Initialize a bit array with all bits set off.
	tracker := bitarray.NewBitArray(26)

	// Iterate over the bytes and see if its a alphabet. If it is mark the array index representing
	// the alphabet as 1 (on).
	for _, v := range []byte(s) {
		index := -1

		switch {
		case v >= 65 && v <= 90:
			index = 90 - int(v)
		case v >= 97 && v <= 122:
			index = 122 - int(v)
		}

		if index != -1 {
			if err := tracker.SetBit(uint64(index % 26)); err != nil {
				fmt.Println("Unable to set bit. Returning false")
				return false
			}
		}
	}

	// Iterate over the bit array and if any of the bits is found to be switched off, return false.
	for i := 0; i < 26; i++ {
		if bit, _ := tracker.GetBit(uint64(i)); !bit {
			return false
		}
	}
	return true
}
