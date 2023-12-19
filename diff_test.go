package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
)

func ExampleDiff() {
	sliceA := []string{"A", "B", "C", "D"}
	sliceB := []string{"A", "G", "B", "F"}
	diff := slt.Diff(sliceA, sliceB)

	fmt.Println(diff)
	// Output:
	// [G F]
}
