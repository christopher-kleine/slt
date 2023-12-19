package slt

// Intersect returns the elements that are in both slices.
func Intersect[S ~[]E, E comparable](a S, b S) S {
	valuesA := make(map[E]*int)
	valuesB := make(map[E]*int)
	result := make(S, 0, min(len(a), len(b)))

	for _, e := range a {
		valuesA[e] = nil
	}

	for _, e := range b {
		_, okA := valuesA[e]
		_, okB := valuesB[e]
		if okA && !okB {
			result = append(result, e)
		}
	}

	return result
}
