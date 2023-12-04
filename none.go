package slt

// None checks if no entry that match the predicate. It's the opposite of [Any]/[Some].
func None[S ~[]E, E any](s S, f func(E) bool) bool {
	for _, e := range s {
		if f(e) {
			return false
		}
	}

	return true
}
