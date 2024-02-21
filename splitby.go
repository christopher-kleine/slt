package slt

// SplitBy splits the slice into 2 parts and returns them.
//
// The first slice is the result of every entry that match the predicate.
//
// The second slice is the result of every entry that failed the predicate.
func SplitBy[S ~[]E, E any](s S, f func(E) bool) (S, S) {
	if f == nil {
		return nil, nil
	}

	left := make(S, 0, len(s))
	right := make(S, 0, len(s))

	for _, v := range s {
		if f(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}
