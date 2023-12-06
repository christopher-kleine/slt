package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
)

func TestMap(t *testing.T) {
	inputList := []int{1, 2, 3}

	square := func(x int) int {
		return x * x
	}

	expected := []int{1, 4, 9}
	actual := slt.Map(inputList, square)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}

func ExampleMap() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	toWord := func(v int) string {
		if v%2 == 0 {
			return "Even"
		} else {
			return "Odd"
		}
	}
	fmt.Println(slt.Map(input, toWord))
	// Output:
	// [Even Odd Even Odd Even Odd Even Odd Even Odd]
}
