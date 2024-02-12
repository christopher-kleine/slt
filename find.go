package slt

// Find returns the first element that match the predicate. It also returns the index.
// In case no element matches, an empty value and -1 are returned.
//
// If the offset is less than zero, it defaults to zero.
func Find[S ~[]E, E any](s S, offset int, f func(E) bool) (E, int) {
	var result E

	if offset < 0 {
		offset = 0
	}
	for index := offset; index < len(s); index++ {
		if f(s[index]) {
			return s[index], index
		}
	}

	return result, -1
}
