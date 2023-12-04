package slt

// Overlap checks if any value in slice A is also found in slice B.
func Overlap[S ~[]E, E comparable](a S, b S) bool {
	values := make(map[E]*int)
	for _, e := range a {
		values[e] = nil
	}

	for _, e := range b {
		if _, ok := values[e]; ok {
			return true
		}
	}

	return false
}

// OverlapFunc checks if any value in slice A is also found in slice B.
func OverlapFunc[S ~[]E, E any](a S, b S, fn func(E, E) bool) bool {
	for _, e1 := range b {
		for _, e2 := range a {
			if fn(e1, e2) {
				return true
			}
		}
	}

	return false
}
