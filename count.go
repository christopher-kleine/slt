package slt

// Count counts the entries that match the predicate
func Count[S ~[]E, E any](s S, fn func(E) bool) int {
	result := 0

	for _, e := range s {
		if fn(e) {
			result++
		}
	}

	return result
}
