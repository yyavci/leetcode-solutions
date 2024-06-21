package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {

	tests := map[string]struct {
		x        int
		expected bool
	}{
		"case 1": {
			x:        121,
			expected: true,
		},
		"case 2": {
			x:        -121,
			expected: false,
		},
		"case 3": {
			x:        10,
			expected: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := isPalindrome(test.x)
			assert.Equal(t, test.expected, res)
		})
	}
}
