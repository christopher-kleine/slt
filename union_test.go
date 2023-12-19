package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
)

func ExampleUnion() {
	a := []string{"A", "B", "C", "D", "H"}
	b := []string{"A", "C", "E", "F", "G"}
	c := []string{"F", "G", "I", "J", "K"}
	d := slt.Union(a, b, c)

	fmt.Println(d)
	// Output:
	// [A B C D H E F G I J K]
}
