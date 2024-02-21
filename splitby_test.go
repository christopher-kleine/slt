package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func ExampleSplitBy() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	even, odd := slt.SplitBy(input, numbers.Even)
	fmt.Println(even)
	fmt.Println(odd)
	// Output:
	// [0 2 4 6 8]
	// [1 3 5 7 9]
}
