package slt

// Diff returns the elements that are only in one of the slices.
func Diff[S ~[]E, E comparable](a S, b S) S {
	valuesA := make(map[E]*int)
	valuesB := make(map[E]*int)
	valuesR := make(map[E]*int)
	result := make(S, 0, len(a)+len(b))

	for _, e := range a {
		valuesA[e] = nil
	}

	for _, e := range b {
		_, okA := valuesA[e]
		_, okB := valuesB[e]
		_, okR := valuesR[e]

		if (okA != !okB) && !okR {
			result = append(result, e)
		}
	}

	return result
}
