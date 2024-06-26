package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInteger(t *testing.T) {

	tests := map[string]struct {
		roman    string
		expected int
	}{
		"case 1": {
			roman:    "III",
			expected: 3,
		},
		"case 2": {
			roman:    "LVIII",
			expected: 58,
		},

		"case 3": {
			roman:    "MCMXCIV",
			expected: 1994,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := romanToInt(test.roman)
			assert.Equal(t, test.expected, res)
		})
	}
}
