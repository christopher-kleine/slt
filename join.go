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
