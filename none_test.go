package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestNone(t *testing.T) {
	inputList := []int{1, 5, 3}

	expected := true
	actual := slt.None(inputList, numbers.Even)
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}
