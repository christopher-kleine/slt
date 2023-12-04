package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestAll(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := false
	actual := slt.All(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}
