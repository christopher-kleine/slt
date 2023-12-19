package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
)

func ExampleIntersect() {
	sliceA := []string{"A", "B", "C", "D"}
	sliceB := []string{"A", "G", "B", "F"}
	diff := slt.Intersect(sliceA, sliceB)

	fmt.Println(diff)
	// Output:
	// [A B]
}
