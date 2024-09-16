package main

import (
	"flag"
	"fmt"
	"github.com/kirkegaard/go-rle/pkg/rle"
	"log"
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
		encoded, err := rle.Encode(*encodeFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(encoded)
	} else if *decodeFlag != "" {
		// Handle decoding
		decoded, err := rle.Decode(*decodeFlag)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(decoded)
	}
}
