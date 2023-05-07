package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"unicode/utf16"

	"golang.design/x/clipboard"
)

func Encode(inputStr string) (string, string) {
	base64EncodedStr := base64.StdEncoding.EncodeToString([]byte(inputStr))

	// Convert LF to CRLF for Windows
	inputStrCRLF := strings.ReplaceAll(inputStr, "\n", "\r\n")
	utf16LEEncodedStr := utf16.Encode([]rune(inputStrCRLF))
	utf16LEBytes := make([]byte, len(utf16LEEncodedStr)*2)
	for i, r := range utf16LEEncodedStr {
		utf16LEBytes[i*2] = byte(r)
		utf16LEBytes[i*2+1] = byte(r >> 8)
	}
	utf16LEBase64EncodedStr := base64.StdEncoding.EncodeToString(utf16LEBytes)

	return base64EncodedStr, utf16LEBase64EncodedStr
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./b64 \"STRING TO ENCODE\"")
		os.Exit(1)
	}

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	inputStr := os.Args[1]
	base64EncodedStr, utf16LEBase64EncodedStr := Encode(inputStr)

	clipboard.Write(clipboard.FmtText, []byte(utf16LEBase64EncodedStr))
	fmt.Printf("Base64 encoded: %s\n", base64EncodedStr)
	fmt.Printf("\nPowerShell encoded (UTF-16LE base64): %s\n", utf16LEBase64EncodedStr)
	fmt.Println("\nPS encoded copied to clipboard!")
}

