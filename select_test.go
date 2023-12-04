package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
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
