package slt

// import "cmp"

// // Mode finds the value that appears the most often in a slice.
// // If there are multiple values with the same amount, it returns the smaller value.
// func Mode[S ~[]E, E cmp.Ordered](s S) (E, int) {
// 	var (
// 		result E
// 		counts = make(map[E]int)
// 		max    = 0
// 	)

// 	for _, v := range s {
// 		counts[v] = counts[v] + 1
// 		if counts[v] >= max {
// 			max = counts[v]
// 			if cmp.Compare(v, result) == -1 {
// 				result = v
// 			}
// 		}
// 	}

// 	return result, max
// }

// // // ModeFunc works the same as [Mode]. But it uses the provided callback.
// // func ModeFunc[S ~[]E, E any, C cmp.Ordered](s S, f func(E) C) (E, int) {
// // 	var (
// // 		result E
// // 		counts = make(map[C]int)
// // 		max    = 0
// // 	)

// // 	for _, v := range s {
// // 		rv := f(v)
// // 		counts[rv] = counts[rv] + 1
// // 		if counts[rv] >= max {
// // 			max = counts[rv]
// // 			if cmp.Compare(rv, result) == -1 {
// // 				result = rv
// // 			}
// // 		}
// // 	}

// // 	return result, max
// // }
