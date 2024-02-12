package slt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/christopher-kleine/slt"
	"github.com/christopher-kleine/slt/numbers"
)

func TestFind(t *testing.T) {
	inputList := []int{1, 2, 3}

	expected := 2
	actual, index := slt.Find(inputList, 0, numbers.Even)
	if index == -1 {
		t.Errorf("expected index to be >=0; got %d instead", index)
	}
	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}

}

func ExampleFind() {
	input := []string{"Anna", "Berta", "chris"}
	isLowerCase := func(s string) bool {
		return strings.ToLower(s) == s
	}
	lowerCaseEntry, index := slt.Find(input, 0, isLowerCase)
	if index > -1 {
		fmt.Println("Found the lower case word entry:", lowerCaseEntry)
	} else {
		fmt.Println("No word is purely lowercase")
	}

	input = []string{"Anna", "Berta", "Chris"}
	lowerCaseEntry, index = slt.Find(input, 0, isLowerCase)
	if index > -1 {
		fmt.Println("Found the lower case word entry:", lowerCaseEntry)
	} else {
		fmt.Println("No word is purely lowercase")
	}
	// Output:
	// Found the lower case word entry: chris
	// No word is purely lowercase
}
