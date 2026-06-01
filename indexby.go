package slt

// IndexBy uses the callback to create a new map based on the callback function.
func IndexBy[S ~[]E, E any, T comparable](s S, fn func(E) T) map[T]E {
	result := make(map[T]E)

	for _, v := range s {
		t := fn(v)
		result[t] = v
	}

	return result
}
