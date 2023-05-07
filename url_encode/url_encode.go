package main

import (
	"fmt"
	"net/url"

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

	// URI encode the reserved characters
	encoded := url.QueryEscape(input)

	// Print the encoded string
	fmt.Println(encoded)

	// Copy the encoded string back to the clipboard
	clipboard.Write(clipboard.FmtText, []byte(encoded))

	// Print confirmation that the encoded string was copied to the clipboard
	fmt.Println("URL Encoded String Copied to Clipboard!")
}

