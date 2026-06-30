package slt

func Lookup[S ~[]C, M ~map[C]A, C comparable, A any](m M, s S) []A {
	result := make([]A, 0, len(m))

	for _, k := range s {
		v, ok := m[k]
		if ok {
			result = append(result, v)
		}
	}

	return result
}
