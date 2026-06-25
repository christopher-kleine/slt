package slt_test

import (
	"fmt"
	"strings"

	"github.com/christopher-kleine/slt/v2"
	"github.com/christopher-kleine/slt/v2/numbers"
)

func ExampleSlice_Map() {
	var (
		input  = []int{9, 3, 6, 10, 2, 1, 8, 5, 7, 4}
		actual = slt.NewSlice(input).Map(func(v int) string {
			return map[int]string{
				1:  "one",
				2:  "two",
				3:  "three",
				4:  "four",
				5:  "five",
				6:  "six",
				7:  "seven",
				8:  "eight",
				9:  "nine",
				10: "ten",
			}[v]
		})
	)

	fmt.Println(actual)

	// Output:
	// [nine three six ten two one eight five seven four]
}

func ExampleSlice_Filter() {
	var (
		input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual = slt.NewSlice(input).Filter(numbers.IsEven)
	)

	fmt.Println(actual)

	// Output:
	// [0 2 4 6 8]
}

func ExampleSlice_Reject() {
	var (
		input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual = slt.NewSlice(input).Reject(numbers.IsEven)
	)

	fmt.Println(actual)

	// Output:
	// [1 3 5 7 9]
}

func ExampleSlice_Unique() {
	var (
		input  = []string{"foo", "bar", "FOO", "BAr", "demo"}
		actual = slt.NewSlice(input).Unique(func(v string) string {
			return strings.ToLower(v)
		})
	)

	fmt.Println(actual)

	// Output:
	// [foo bar demo]
}
