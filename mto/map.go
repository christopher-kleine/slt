package mto

func Map[M ~map[K]E, E any, K comparable, T any](m M, fn func(K, E) T) map[K]T {
	result := make(map[K]T)

	for k, v := range m {
		result[k] = fn(k, v)
	}

	return result
}
