package slt

// Any checks if there is at least one entry that match the predicate.
func Any[S ~[]E, E any](s S, fn func(E) bool) bool {
	for _, e := range s {
		if fn(e) {
			return true
		}
	}

	return false
}

// Some is an alias for [Any].
func Some[S ~[]E, E any](s S, fn func(E) bool) bool {
	return Any(s, fn)
}
