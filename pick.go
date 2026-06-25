package slt

// import (
// 	"crypto/rand"
// 	"io"
// 	"math/big"
// )

// // Pick selects random entries from s. It can contain duplicates.
// //
// // It panics if len(s) == 0.
// //
// // If n > len(s), the result will contain duplicates for certain.
// //
// // if n <= 0, n is assumed to be len(n).
// func Pick[S ~[]E, E any](s S, n int, r io.Reader) S {
// 	result := make(S, n)

// 	if n <= 0 {
// 		n = len(s)
// 	}

// 	max := big.NewInt(int64(len(s)))

// 	for index := range n {
// 		i, err := rand.Int(r, max)
// 		if err != nil {
// 			continue
// 		}

// 		si := int(i.Int64())
// 		result[index] = s[si]
// 	}

// 	return result
// }

// // PickUnique works similiar to [Pick].
// // But it picks the entries using Fisher-Yates shuffle.
// //
// // It panics if len(n) == 0.
// //
// // If n > len(s) or n <= 0, n is assumed to be len(n).
// func PickUnique[S ~[]E, E any](s S, n int, r io.Reader) S {
// 	cp := s[:]

// 	if n <= 0 {
// 		n = len(s)
// 	}

// 	if n > len(s) {
// 		n = len(s)
// 	}

// 	for index := range cp {
// 		max := big.NewInt(int64(index + 1))
// 		i, err := rand.Int(r, max)
// 		if err != nil {
// 			continue
// 		}

// 		si := int(i.Int64())

// 		cp[index], cp[si] = cp[si], cp[index]

// 		if index >= n {
// 			break
// 		}
// 	}

// 	return cp[:n]
// }
