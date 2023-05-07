package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

func getAbsolutePath(fileName string) (string, error) {
	return filepath.Abs(fileName)
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument.")
		os.Exit(1)
	}

	fileName := os.Args[1]
	absPath, err := getAbsolutePath(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("\nPath copied to clipboard!\n", absPath)
	clipboard.Write(clipboard.FmtText, []byte(absPath))
}
