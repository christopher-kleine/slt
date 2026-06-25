package slt_test

// import (
// 	"fmt"
// 	"math/rand"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/christopher-kleine/slt/v2"
// )

// const maxSize = 10000

// func TestUnique(t *testing.T) {
// 	var (
// 		input    = []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
// 		expected = []int{0, 1, 2, 3, 4, 5, 6}
// 		actual   = slt.Unique(input)
// 	)

// 	assert.Equal(t, expected, actual)
// }

// func TestUniqueFunc(t *testing.T) {
// 	type person struct {
// 		Name string
// 		Age  int
// 	}

// 	var (
// 		hero     = person{Name: "Hero", Age: 20}
// 		villain  = person{Name: "Villain", Age: 30}
// 		sidekick = person{Name: "Sidekick", Age: 25}
// 		input    = []person{hero, villain, hero, sidekick, villain}
// 		expected = []person{hero, villain, sidekick}
// 		actual   = slt.UniqueFunc(input, func(a person) string {
// 			return a.Name
// 		})
// 	)

// 	assert.Equal(t, expected, actual)
// }

// func ExampleUnique() {
// 	input := []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
// 	fmt.Println(slt.Unique(input))
// 	// Output:
// 	// [0 1 2 3 4 5 6]
// }

// func TestUniqueUnstable(t *testing.T) {
// 	var (
// 		input    = []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
// 		expected = []int{0, 1, 2, 3, 4, 5, 6}
// 		actual   = slt.UniqueUnstable(input)
// 	)

// 	assert.ElementsMatch(t, expected, actual)
// }

// func BenchmarkUnique(b *testing.B) {
// 	var (
// 		input = []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
// 	)

// 	for range b.N {
// 		_ = slt.Unique(input)
// 	}
// }

// func BenchmarkUniqueUnstable(b *testing.B) {
// 	var (
// 		input = []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
// 	)

// 	for range b.N {
// 		_ = slt.UniqueUnstable(input)
// 	}
// }

// func generateData(n int) []int {
// 	data := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		data[i] = rand.Intn(n / 2) // intentionally many duplicates
// 	}
// 	return data
// }

// func BenchmarkUniqueSizes(b *testing.B) {
// 	sizes := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}

// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("Stable_%d", size), func(b *testing.B) {
// 			data := generateData(size)
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_ = slt.Unique(data)
// 			}
// 		})

// 		b.Run(fmt.Sprintf("Unstable_%d", size), func(b *testing.B) {
// 			data := generateData(size)
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_ = slt.UniqueUnstable(data)
// 			}
// 		})
// 	}
// }
