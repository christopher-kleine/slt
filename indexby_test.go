package slt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christopher-kleine/slt"
)

func TestIndexByString(t *testing.T) {
	testTable := []struct {
		values []string
		want   map[string]string
		fn     func(string) string
	}{
		{
			values: []string{"Anna", "Bob", "Carmen", "Domenik", "Chris"},
			want: map[string]string{
				"a": "Anna",
				"b": "Bob",
				"c": "Chris",
				"d": "Domenik",
			},
			fn: func(s string) string {
				return string(strings.ToLower(s))[:1]
			},
		},
	}

	for _, testCase := range testTable {
		have := slt.IndexBy(testCase.values, testCase.fn)
		want := testCase.want

		assert.Equal(t, have, want)
	}
}

func TestIndexByInt(t *testing.T) {
	testTable := []struct {
		values []int
		want   map[int]int
		fn     func(int) int
	}{
		{
			values: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: map[int]int{
				0: 8,
				1: 9,
			},
			fn: func(v int) int {
				return v % 2
			},
		},
	}

	for _, testCase := range testTable {
		have := slt.IndexBy(testCase.values, testCase.fn)
		want := testCase.want

		assert.Equal(t, have, want)
	}
}

func ExampleIndexBy() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	toWord := func(v int) string {
		if v%2 == 0 {
			return "Even"
		} else {
			return "Odd"
		}
	}
	fmt.Println(slt.IndexBy(input, toWord))
	// Output:
	// map[Even:8 Odd:9]
}
