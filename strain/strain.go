// Package strain provides functionality to identify the keeping elements
// and discarding elements from any type of supported collections.
package strain

// Ints represents the list of integers
type Ints []int

// Strings represents the list of strings
type Strings []string

// Lists represents a list of list of integers
type Lists [][]int

// Keep returns Ints where the given predicate function is true.
func (ints Ints) Keep(fn func(int) bool) Ints {
	var keep Ints
	for _, v := range ints {
		if fn(v) {
			keep = append(keep, v)
		}
	}

	return keep
}

// Discard returns Ints where the given predicate function is false.
func (ints Ints) Discard(fn func(int) bool) Ints {
	var discard Ints
	for _, v := range ints {
		if !fn(v) {
			discard = append(discard, v)
		}
	}

	return discard
}

// Keep returns Lists where the given predicate function is true.
func (lists Lists) Keep(fn func([]int) bool) Lists {
	var keep Lists
	for _, v := range lists {
		if fn(v) {
			keep = append(keep, v)
		}
	}

	return keep
}

// Keep returns Strings where the given predicate function is true.
func (strings Strings) Keep(fn func(string) bool) Strings {
	var keep Strings
	for _, v := range strings {
		if fn(v) {
			keep = append(keep, v)
		}
	}

	return keep
}
