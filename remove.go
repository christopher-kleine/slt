package slt

// Remove filters the array based on the predicate.
func Remove[S ~[]E, E any](s S, f func(E) bool) S {
	result := make(S, 0, len(s))

	for _, e := range s {
		if !f(e) {
			result = append(result, e)
		}
	}

	return result
}
