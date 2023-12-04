package slt_test

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christopher-kleine/slt"
)

func TestUnique(t *testing.T) {
	var (
		input    = []int{0, 1, 2, 3, 4, 5, 6, 5, 1}
		expected = []int{0, 1, 2, 3, 4, 5, 6}
		actual   = slt.Unique(input)
	)

	assert.Equal(t, expected, actual)
}

func TestUniqueFunc(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	var (
		hero     = person{Name: "Hero", Age: 20}
		villain  = person{Name: "Villain", Age: 30}
		sidekick = person{Name: "Sidekick", Age: 25}
		input    = []person{hero, villain, hero, sidekick, villain}
		expected = []person{hero, sidekick, villain}
		actual   = slt.UniqueFunc(input, func(a person, b person) int {
			return cmp.Compare(a.Name, b.Name)
		})
	)

	assert.Equal(t, expected, actual)
}
