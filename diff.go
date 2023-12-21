package slt

// Diff returns the elements that are only in one of the slices.
func Diff[S ~[]E, E comparable](a S, b S) S {
	var (
		valuesA = make(map[E]*int)
		valuesB = make(map[E]*int)
		valuesR = make(map[E]*int)
		result  = make(S, 0, len(a)+len(b))
	)

	for _, e := range a {
		valuesA[e] = nil
	}

	for _, e := range b {
		valuesB[e] = nil
	}

	for _, e := range a {
		_, okB := valuesB[e]
		_, okR := valuesR[e]

		if okB || okR {
			continue
		}

		result = append(result, e)
		valuesR[e] = nil
	}

	for _, e := range b {
		_, okA := valuesA[e]
		_, okR := valuesR[e]

		if okA || okR {
			continue
		}

		result = append(result, e)
		valuesR[e] = nil
	}

	return result
}
