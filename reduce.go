package slt

// Reduce reduces a collection using the function fn
func Reduce[S ~[]E, E any](s S, start E, f func(acc E, curr E) E) E {
	result := start

	for _, e := range s {
		result = f(result, e)
	}

	return result
}
