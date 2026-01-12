package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <fileName>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	fmt.Println("File name:", fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Create a buffer to store the file contents
	var contents bytes.Buffer

	// Copy file contents to the buffer
	_, err = io.Copy(&contents, file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Print the file contents
	fmt.Println("\nFile contents:")
	fmt.Println(contents.String())
}
