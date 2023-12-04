package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestCount(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := 2
	actual := slt.Count(inputList, numbers.Odd)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}
