package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestNone(t *testing.T) {
	inputList := []int{1, 5, 3}

	expected := true
	actual := slt.None(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func ExampleNone() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slt.None(input, numbers.Even))
	fmt.Println(slt.None(input, numbers.Below(10, true)))
	// Output:
	// false
	// false
}
