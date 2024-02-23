package slt

import "slices"

// Equal is a modified version of [slices.Equal]. It compares two slices on their content.
//
// Is "inOrder" false, the order of the content doesn't matter. If it's true,
// this function becomes an alias tp [slices.Equal].
func Equal[S ~[]E, E comparable](s1 S, s2 S, inOrder bool) bool {
	if inOrder {
		return slices.Equal(s1, s2)
	}

	if len(s1) != len(s2) {
		return false
	}

	lookup := make(map[E]bool)
	for _, v := range s1 {
		lookup[v] = true
	}
	for _, v := range s2 {
		if _, ok := lookup[v]; !ok {
			return false
		}
	}

	return true
}
