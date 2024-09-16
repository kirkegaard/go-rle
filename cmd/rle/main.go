package main

import (
	"flag"
	"fmt"
	"github.com/kirkegaard/go-rle/pkg/rle"
)

func main() {
	// Define command-line flags
	encodeFlag := flag.String("encode", "", "String to encode")
	decodeFlag := flag.String("decode", "", "String to decode")

	// Parse the command-line flags
	flag.Parse()

	// Check if both encode and decode are provided or neither
	if (*encodeFlag != "" && *decodeFlag != "") || (*encodeFlag == "" && *decodeFlag == "") {
		fmt.Println("Please provide either --encode or --decode flag.")
		flag.Usage()
		return
	}

	if *encodeFlag != "" {
		// Handle encoding
		encoded := rle.Encode(*encodeFlag)
		fmt.Println(encoded)
	} else if *decodeFlag != "" {
		// Handle decoding
		decoded := rle.Decode(*decodeFlag)
		fmt.Println(decoded)
	}
}
