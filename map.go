package slt

// Map applies a transform function on all entries in a slice and returns the new slice.
func Map[S ~[]E, E any, T any](s S, fn func(E) T) []T {
	result := make([]T, len(s))

	for k, v := range s {
		result[k] = fn(v)
	}

	return result
}
