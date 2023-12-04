package slt

import "slices"

// Unique makes sure every entry only appears once. It won't change the order of entries.
func Unique[S ~[]E, E comparable](s S) S {
	entryMap := make(map[E]*int)
	result := make(S, 0, len(s))

	for _, e := range s {
		if _, ok := entryMap[e]; !ok {
			result = append(result, e)
			entryMap[e] = nil
		}
	}

	return result
}

// UniqueFunc makes sure every entry only appears once. Unlike [Unique] this changes the order of entries.
func UniqueFunc[S ~[]E, E any](s S, cmp func(a, b E) int) S {
	result := slices.Clone(s)
	slices.SortFunc(result, cmp)
	result = slices.CompactFunc(result, func(a, b E) bool {
		return cmp(a, b) == 0
	})

	return result
}
