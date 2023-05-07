package main

import (
	"bytes"
	"fmt"

	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	// Read text data from the clipboard
	data := clipboard.Read(clipboard.FmtText)

	// Convert byte data to string
	input := string(data)

	// Escape double quotes, except if they are already escaped
	var buf bytes.Buffer
	for i := 0; i < len(input); i++ {
		if input[i] == '"' && (i == 0 || input[i-1] != '\\') {
			buf.WriteString(`\"`)
		} else {
			buf.WriteByte(input[i])
		}
	}

	escaped := buf.String()

	// Print the escaped string
	fmt.Println(escaped)

	// Copy the escaped string back to the clipboard
	clipboard.Write(clipboard.FmtText, []byte(escaped))
}

