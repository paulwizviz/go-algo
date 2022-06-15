package roman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntSlice(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			value []int
			err   error
		}
		description string
	}{
		{
			input: "IV",
			expected: struct {
				value []int
				err   error
			}{
				value: []int{I, V},
				err:   nil,
			},
			description: "Valid sequence",
		},
		{
			input: "iV",
			expected: struct {
				value []int
				err   error
			}{
				value: nil,
				err:   Error,
			},
			description: "Invalid case",
		},
		{
			input: "I2",
			expected: struct {
				value []int
				err   error
			}{
				value: nil,
				err:   Error,
			},
			description: "Invalid character",
		},
	}

	for index, test := range tests {
		actual, actualErr := stringToIntSlice(test.input)

		assert.Equal(t, test.expected.value, actual, fmt.Sprintf("Test: %d.1 Description: %s", index, test.description))
		assert.Equal(t, test.expected.err, actualErr, fmt.Sprintf("Test: %d.2 Description: %s", index, test.description))
	}
}

func TestSymbolsToValue(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		description string
	}{
		{input: "MCMLXXVI", expected: 1976, description: "correct"},
		{input: "IV", expected: 4, description: "smaller"},
		{input: "MMXVIII", expected: 2018, description: "correct"},
		{input: "MCMLXV", expected: 1965, description: "correct"},
	}

	for index, test := range tests {

		actual, _ := SymbolsToValue(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("Test: %d.1 Description: %s", index, test.description))

	}
}
