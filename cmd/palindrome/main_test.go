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
		},
		{
			input:       "redrosesrunnorisksironnursesorder",
			expected:    true,
			description: "Red roses ...",
		},
		{
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

func TestIsPalindromeC(t *testing.T) {
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
		},
		{
			input:       "redrosesrunnorisksironnursesorder",
			expected:    true,
			description: "Red roses ...",
		},
		{
			input:       "selfish",
			expected:    false,
			description: "selfish",
		},
	}

	for index, test := range tests {
		actual := isPalindromeAlt(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("Test: %d Description: %s", index, test.description))
	}
}

var bInput = "redrosesrunnorisksironnursesorder"

func BenchmarkIsIsPalindrome(b *testing.B) {
	for count := 0; count < b.N; count++ {
		isPalindrome(bInput)
	}
}

func BenchmarkIsIsPalindromeC(b *testing.B) {
	for count := 0; count < b.N; count++ {
		isPalindromeAlt(bInput)
	}
}
