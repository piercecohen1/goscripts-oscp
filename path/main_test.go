package main

import (
	"path/filepath"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	fileName := "test.txt"

	absPath, err := getAbsolutePath(fileName)
	expected, _ := filepath.Abs(fileName)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if absPath != expected {
		t.Errorf("Expected %s, but got %s", expected, absPath)
	}
}
