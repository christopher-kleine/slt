package slt

// Group uses the callback to group the provided slice into a map of slices.
func Group[S ~[]E, E any, T comparable](s S, fn func(E) T) map[T]S {
	result := make(map[T]S)

	for _, v := range s {
		t := fn(v)
		if _, ok := result[t]; !ok {
			result[t] = make(S, 0, len(s))
		}
		result[t] = append(result[t], v)
	}

	return result
}
