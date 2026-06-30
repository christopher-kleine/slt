package slt

// Join takes 2 slices and combines them based on the returned key on both sides.
func Join[SA ~[]EA, SB ~[]EB, EA any, EB any, JoinKey comparable, ResultKey comparable](a SA, b SB, akey func(EA) JoinKey, bkey func(EB) JoinKey, resultKey func(EA) ResultKey) map[ResultKey]EB {
	index := make(map[JoinKey]EB, len(b))

	for _, e := range b {
		index[bkey(e)] = e
	}

	result := make(map[ResultKey]EB, len(a))

	for _, e := range a {
		if v, ok := index[akey(e)]; ok {
			result[resultKey(e)] = v
		}
	}

	return result
}

// InnerJoin takes a slice and a lookup map. It then returns
func InnerJoin[M ~map[T]E1, S ~[]E2, T comparable, E1 any, E2 any](slice S, lookup M, fn func(E2) T) []E1 {
	result := make([]E1, 0, len(slice))

	for _, v := range slice {
		k := fn(v)
		if v2, ok := lookup[k]; ok {
			result = append(result, v2)
		}
	}

	return result
}

func LeftJoin[M ~map[T]E1, S ~[]E2, T comparable, E1 any, E2 any](slice S, lookup M, fn func(E2) T) []struct {
	Value E1
	OK    bool
} {
	result := make([]struct {
		Value E1
		OK    bool
	}, len(slice))

	for index, v := range slice {
		k := fn(v)
		v2, ok := lookup[k]
		result[index] = struct {
			Value E1
			OK    bool
		}{
			Value: v2,
			OK:    ok,
		}
	}

	return result
}
