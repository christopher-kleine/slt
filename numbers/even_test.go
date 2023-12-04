package numbers_test

import (
	"testing"

	"github.com/christopher-kleine/slt/numbers"
)

func TestEven(t *testing.T) {
	if numbers.Even(1) != false {
		t.Errorf("expected %v != actual %v", false, true)
	}

	if numbers.Even(2) != true {
		t.Errorf("expected %v != actual %v", true, false)
	}
}
