package slt

// Union takes a number of slices and combines them into one.
// The difference between Union and the built-in function append lies in the removal of duplicates.
func Union[S ~[]E, E comparable](s ...S) S {
	lookup := make(map[E]bool)
	sum := 0

	for _, e := range s {
		sum += len(e)
	}

	result := make(S, 0, sum)
	for _, e := range s {
		for _, v := range e {
			if _, ok := lookup[v]; !ok {
				lookup[v] = true
				result = append(result, v)
			}
		}
	}

	return result
}
