package slt_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christopher-kleine/slt"
)

func TestGroupByString(t *testing.T) {
	testTable := []struct {
		values []string
		want   map[string][]string
		fn     func(string) string
	}{
		{
			values: []string{"Anna", "Bob", "Carmen", "Domenik", "Chris"},
			want: map[string][]string{
				"a": {"Anna"},
				"b": {"Bob"},
				"c": {"Carmen", "Chris"},
				"d": {"Domenik"},
			},
			fn: func(s string) string {
				return string(strings.ToLower(s))[:1]
			},
		},
	}

	for _, testCase := range testTable {
		have := slt.Group(testCase.values, testCase.fn)
		want := testCase.want

		assert.Equal(t, have, want)
	}
}

func TestGroupByInt(t *testing.T) {
	testTable := []struct {
		values []int
		want   map[int][]int
		fn     func(int) int
	}{
		{
			values: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: map[int][]int{
				0: {0, 2, 4, 6, 8},
				1: {1, 3, 5, 7, 9},
			},
			fn: func(v int) int {
				return v % 2
			},
		},
	}

	for _, testCase := range testTable {
		have := slt.Group(testCase.values, testCase.fn)
		want := testCase.want

		assert.Equal(t, have, want)
	}
}
