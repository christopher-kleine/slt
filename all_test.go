package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestAll(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := false
	actual := slt.All(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func ExampleAll() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slt.All(input, numbers.Even))
	fmt.Println(slt.All(input, numbers.Below(10, true)))
	// Output:
	// false
	// true
}
