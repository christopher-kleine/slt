package slt_test

import (
	"fmt"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestAny(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := true
	actual := slt.Any(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func TestAny2(t *testing.T) {
	inputList := []int{1, 5, 3}

	expected := false
	actual := slt.Any(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func ExampleAny() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slt.Any(input, numbers.Even))
	fmt.Println(slt.All(input, numbers.Above(10, true)))
	// Output:
	// true
	// false
}
