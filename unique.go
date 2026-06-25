package slt

// // Unique makes sure every entry only appears once. It won't change the order of entries.
// func Unique[S ~[]E, E comparable](s S) S {
// 	seen := make(map[E]struct{})
// 	result := make(S, 0, len(s))

// 	for _, e := range s {
// 		if _, ok := seen[e]; !ok {
// 			result = append(result, e)
// 			seen[e] = struct{}{}
// 		}
// 	}

// 	return result
// }

// // UniqueFunc makes sure every entry only appears once. It won't change the order of entries.
// func UniqueFunc[S ~[]E, E any, C comparable](s S, cmp func(e E) C) S {
// 	entryMap := make(map[C]*int)
// 	result := make(S, 0, len(s))

// 	for _, e := range s {
// 		c := cmp(e)
// 		if _, ok := entryMap[c]; !ok {
// 			result = append(result, e)
// 			entryMap[c] = nil
// 		}
// 	}

// 	return result
// }

// // UniqueUnstable returns a slice with duplicate elements removed.
// // The order of elements in the result is not guaranteed.
// func UniqueUnstable[S ~[]E, E comparable](s S) S {
// 	seen := make(map[E]struct{}, len(s))
// 	for _, e := range s {
// 		seen[e] = struct{}{}
// 	}

// 	result := make(S, 0, len(seen))
// 	for e := range seen {
// 		result = append(result, e)
// 	}

// 	return result
// }
