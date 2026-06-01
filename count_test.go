package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestCount(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := 2
	actual := slt.Count(inputList, numbers.IsOdd)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func ExampleCount() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slt.Count(input, numbers.IsEven))
	fmt.Println(slt.Count(input, numbers.AboveOrEqual(8)))
	// Output:
	// 5
	// 2
}
