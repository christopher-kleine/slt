package slt_test

import (
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
