package slt

import "slices"

func ReplaceValues[S ~[]E, M map[E]E, E comparable](s S, m M) S {
	if len(m) == 0 || len(s) == 0 {
		return s
	}

	var (
		result S
		cloned = false
	)

	for i, e := range s {
		if v, ok := m[e]; ok {
			if !cloned {
				result = slices.Clone(s)
				cloned = true
			}
			result[i] = v
		}
	}

	if cloned {
		return result
	}

	return s
}
