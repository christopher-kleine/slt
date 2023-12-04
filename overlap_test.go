package slt_test

import (
	"testing"

	"github.com/christopher-kleine/slt"
)

func TestOverlap(t *testing.T) {
	ListA := []int{1, 2, 3}
	ListB := []int{1, 5, 7}
	ListC := []int{5, 7, 9}

	Found := slt.Overlap(ListA, ListB)
	if Found == false {
		t.Errorf("expected %v != actual %v", true, Found)
	}

	NotFound := slt.Overlap(ListA, ListC)
	if NotFound == true {
		t.Errorf("expected %v != actual %v", false, NotFound)
	}
}
