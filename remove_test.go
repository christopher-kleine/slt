package slt_test

import (
	"fmt"

	"github.com/christopher-kleine/slt"
)

func ExampleRemove() {
	input := []string{"H", "E", "L", "L", "O", "W", "O", "R", "L", "D"}
	removed := slt.Remove(input, "A", "E", "I", "O", "U")
	fmt.Println(removed)
	// Output: [H L L W R L D]
}
