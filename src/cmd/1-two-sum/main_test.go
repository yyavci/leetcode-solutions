package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	tests := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"case 1": {
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		"case 2": {
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},

		"case 3": {
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := twoSum(test.nums, test.target)
			assert.Equal(t, test.expected, res)
		})
	}
}
