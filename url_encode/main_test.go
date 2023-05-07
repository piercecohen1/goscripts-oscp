package main

import (
	"testing"
)

func TestUrlEncode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "This is a test string with special characters: !@#$%^&*()",
			expected: "This+is+a+test+string+with+special+characters%3A+%21%40%23%24%25%5E%26%2A%28%29",
		},
		{
			input:    "https://www.example.com/search?q=hello world",
			expected: "https%3A%2F%2Fwww.example.com%2Fsearch%3Fq%3Dhello+world",
		},
		{
			input:    "Encode spaces:   and tabs:\t",
			expected: "Encode+spaces%3A+++and+tabs%3A%09",
		},
		{
			input:    "Test & More",
			expected: "Test+%26+More",
		},
		{
			input:    "Slash / Backslash \\",
			expected: "Slash+%2F+Backslash+%5C",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			encoded := urlEncode(tc.input)

			if encoded != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, encoded)
			}
		})
	}
}
