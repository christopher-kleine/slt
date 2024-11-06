package slt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/christopher-kleine/slt"
)

func TestReduce(t *testing.T) {
	inputList := []int{1, 2, 3}

	funcReduce := func(acc int, curr int) int {
		return curr + acc
	}

	expected := 6
	accumulator := 0
	actual := slt.Reduce(inputList, accumulator, funcReduce)

	if expected != actual {
		t.Errorf("expected %v != actual %v", expected, actual)
	}
}

func ExampleReduce() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add := func(acc int, curr int) int {
		return acc + curr
	}
	fmt.Println("Sum of all numbers between 1 and 10 is:", slt.Reduce(numbers, 0, add))
	// Output:
	// Sum of all numbers between 1 and 10 is: 55
}

func ExampleReduce2() {
	numbers := []string{
		"A", "b", "C", "D", "E", "f", "g", "H", "I", "J", "K", "L", "M",
		"n", "O", "p", "q", "r", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	addLower := func(acc int, curr string) int {
		if curr == strings.ToLower(curr) {
			return acc + 1
		}

		return acc
	}
	fmt.Println("Sum of all lower letters is:", slt.Reduce(numbers, 0, addLower))
	// Output:
	// Sum of all lower letters is: 7
}
