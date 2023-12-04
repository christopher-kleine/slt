package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
)

func TestReduce(t *testing.T) {
	inputList := []int{1, 2, 3}

	funcReduce := func(acc int, curr int) int {
		return curr + acc
	}

	expected := 6
	accumulator := 0
	actual := slt.Reduce(inputList, accumulator, funcReduce)

	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}
