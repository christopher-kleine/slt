package slt

// All checks if every entry in the slice match the predicate.
func All[S ~[]E, E any](s S, fn func(E) bool) bool {
	for _, e := range s {
		if !fn(e) {
			return false
		}
	}

	return true
}

// Every is an alias for [All].
func Every[S ~[]E, E any](s S, fn func(E) bool) bool {
	return All(s, fn)
}
