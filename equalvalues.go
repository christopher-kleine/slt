package slt

// // EqualValues compares two slices on their content. The order of the elements
// // don't matter.
// func EqualValues[S ~[]E, E comparable](s1 S, s2 S) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}

// 	lookup := make(map[E]bool)
// 	for _, v := range s1 {
// 		lookup[v] = true
// 	}
// 	for _, v := range s2 {
// 		if _, ok := lookup[v]; !ok {
// 			return false
// 		}
// 	}

// 	return true
// }
