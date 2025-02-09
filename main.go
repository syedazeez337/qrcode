package main

import (
	"flag"
	"fmt"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// Define command-line flags for input text, output filename, and size.
	text := flag.String("text", "", "Text to encode into QR code")
	output := flag.String("output", "qr.png", "Output PNG filename")
	size := flag.Int("size", 256, "Size (in pixels) of the generated QR code")
	flag.Parse()

	// Check if text was provided.
	if *text == "" {
		fmt.Println("Usage: go run main.go -text \"your text here\" [-output filename] [-size pixels]")
		return
	}

	// Generate the QR code and save it as a PNG file.
	err := qrcode.WriteFile(*text, qrcode.Medium, *size, *output)
	if err != nil {
		log.Fatalf("Error generating QR code: %v", err)
	}

	fmt.Printf("QR code successfully generated and saved to %s\n", *output)
}
