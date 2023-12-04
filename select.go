package slt

// Select filters the given slice by using the predicate.
func Select[S ~[]E, E any](s S, f func(E) bool) S {
	result := make(S, 0, len(s))

	for _, e := range s {
		if f(e) {
			result = append(result, e)
		}
	}

	return result
}
