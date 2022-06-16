package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "anna",
			expected:    true,
			description: "anna",
		},
		{
			input:       "racecar",
			expected:    true,
			description: "racecar",
		}, {
			input:       "selfish",
			expected:    false,
			description: "selfish",
		},
	}

	for index, test := range tests {
		actual := isPalindrome(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("Test: %d Description: %s", index, test.description))
	}
}
