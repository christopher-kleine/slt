package slt

// Reject returns a new slice, removing all elements using the predicate.
func Reject[S ~[]E, E any](s S, f func(E) bool) S {
	result := make(S, 0, len(s))

	for _, e := range s {
		if !f(e) {
			result = append(result, e)
		}
	}

	return result
}
