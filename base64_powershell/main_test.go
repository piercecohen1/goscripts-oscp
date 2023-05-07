package main

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"
)

func normalizeNewline(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\r", "\n")
}

func TestEncoding(t *testing.T) {
	testCases := []struct {
		input              string
		expectedBase64     string
		expectedPowerShell string
	}{
		{
			input:              "Hello, World!",
			expectedBase64:     "SGVsbG8sIFdvcmxkIQ==",
			expectedPowerShell: "SABlAGwAbABvACwAIABXAG8AcgBsAGQAIQA=",
		},
		{
			input:              "Golang",
			expectedBase64:     "R29sYW5n",
			expectedPowerShell: "RwBvAGwAYQBuAGcA",
		},
		{
			input:              "¡Hola, mundo!",
			expectedBase64:     "wqFIb2xhLCBtdW5kbyE=",
			expectedPowerShell: "oQBIAG8AbABhACwAIABtAHUAbgBkAG8AIQA=",
		},
		{
			input:              "こんにちは世界",
			expectedBase64:     "44GT44KT44Gr44Gh44Gv5LiW55WM",
			expectedPowerShell: "UzCTMGswYTBvMBZOTHU=",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			base64Encoded, powerShellEncoded := Encode(tc.input)

			if base64Encoded != tc.expectedBase64 {
				t.Errorf("Expected %s, but got %s", tc.expectedBase64, base64Encoded)
			}

			expectedPowerShellBytes, _ := base64.StdEncoding.DecodeString(tc.expectedPowerShell)
			powerShellEncodedBytes, _ := base64.StdEncoding.DecodeString(powerShellEncoded)

			if !bytes.Equal(expectedPowerShellBytes, powerShellEncodedBytes) {
				t.Errorf("Expected %s, but got %s", tc.expectedPowerShell, powerShellEncoded)
			}
		})
	}
}
