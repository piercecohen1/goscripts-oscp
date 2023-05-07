package main

import (
	"testing"
)

func TestEscapeDoubleQuotes(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    `This is a "test" string with "quotes".`,
			expected: `This is a \"test\" string with \"quotes\".`,
		},
		{
			input:    `Escaped "quotes" and \"quotes\" mixed.`,
			expected: `Escaped \"quotes\" and \"quotes\" mixed.`,
		},
		{
			input:    `""`,
			expected: `\"\"`,
		},
		{
			input:    `No quotes here.`,
			expected: `No quotes here.`,
		},
		{
			input:    `\"`,
			expected: `\"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			escaped := escapeDoubleQuotes(tc.input)
			if escaped != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, escaped)
			}
		})
	}
}
