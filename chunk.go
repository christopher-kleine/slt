package slt

// type SeparatorHandling string

// const (
// 	// SeparatorDiscard discards the separator. It won't be part of the chunks.
// 	SeparatorDiscard SeparatorHandling = "discard"

// 	// SeparatorSameChunk keeps the separator in the same chunk and starts a
// 	// fresh chunk afterwards.
// 	SeparatorSameChunk SeparatorHandling = "same"

// 	// SeparatorNextChunk places the separator in the next chunk. This may produce
// 	// an additional chunk with only the separator.
// 	SeparatorNextChunk SeparatorHandling = "next"
// )

// func Chunk[S ~[]E, E any](s S, size int) []S {
// 	if size <= 0 {
// 		panic("chunk size must be greater than 0")
// 	}

// 	var result []S
// 	for i := 0; i < len(s); i += size {
// 		end := i + size

// 		if end > len(s) {
// 			end = len(s)
// 		}

// 		result = append(result, s[i:end])
// 	}

// 	return result
// }

// // Chunk splits a Slice based on the provided callback. If the callback returns
// // true, a new chunk will be added.
// func ChunkBy[S ~[]E, E any](s S, separatorHandling SeparatorHandling, fn func(e E) bool) []S {
// 	var (
// 		result []S
// 		chunk  = make(S, 0, 100)
// 	)

// 	for _, e := range s {
// 		if fn(e) {
// 			if separatorHandling == SeparatorSameChunk {
// 				chunk = append(chunk, e)
// 			}

// 			result = append(result, chunk)
// 			chunk = make(S, 0, 100)

// 			if separatorHandling == SeparatorNextChunk {
// 				chunk = append(chunk, e)
// 			}
// 		} else {
// 			chunk = append(chunk, e)
// 		}
// 	}

// 	if len(chunk) > 0 {
// 		result = append(result, chunk)
// 	}

// 	return result
// }
