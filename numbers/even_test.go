package numbers_test

import (
	"testing"

	"github.com/christopher-kleine/slt/numbers"
)

func TestEven(t *testing.T) {
	if numbers.IsEven(1) != false {
		t.Errorf("expected %v != actual %v", false, true)
	}

	if numbers.IsEven(2) != true {
		t.Errorf("expected %v != actual %v", true, false)
	}
}
