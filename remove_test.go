package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestRemove(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := []int{1, 3}
	actual := slt.Remove(inputList, numbers.Even)

	for i, e := range expected {
		if e != actual[i] {
			t.Errorf("expected %v != actual %v", expected, actual)
		}
	}
}
