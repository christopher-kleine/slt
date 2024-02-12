package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestReject(t *testing.T) {
	inputList := []int{1, 2, 3}

	modulotwo := func(x int) bool {
		return x%2 == 0
	}

	expected := []int{1, 3}
	actual := slt.Reject(inputList, modulotwo)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}

func ExampleReject() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Even Numbers:", slt.Reject(input, numbers.Even))
	// Output:
	// Even Numbers: [1 3 5 7 9]
}
