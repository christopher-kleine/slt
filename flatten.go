package slt

func Flatten[S ~[]E, E any](s []S) S {
	var totalLen int

	for _, inner := range s {
		totalLen += len(inner)
	}

	result := make(S, 0, totalLen)

	for _, inner := range s {
		result = append(result, inner...)
	}

	return result
}
