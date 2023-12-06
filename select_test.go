package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestSelect(t *testing.T) {
	inputList := []int{1, 2, 3}

	modulotwo := func(x int) bool {
		return x%2 == 0
	}

	expected := []int{2}
	actual := slt.Select(inputList, modulotwo)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}

func ExampleSelect() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Even Numbers:", slt.Select(input, numbers.Even))
	// Output:
	// Even Numbers: [0 2 4 6 8]
}
