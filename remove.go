package slt

// Remove values from a slice.
func Remove[S ~[]E, E comparable](s S, e ...E) S {
	result := make(S, 0, len(s))

	existMap := make(map[E]byte)
	for _, v := range e {
		existMap[v] = 0
	}

	for _, v := range s {
		if _, ok := existMap[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}
