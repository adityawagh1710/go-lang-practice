package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create file
	file, err := os.Create("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Hi... there")

	if err != nil {
		log.Fatal(err)
	}

	// IMPORTANT: Close before reopening
	file.Close()

	// Open file for reading
	file, err = os.Open("file.txt")

	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}

	defer file.Close()

	// Read file
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("Cannot read file: %v", err)
	}

	fmt.Println(string(data))

	// File stats
	info, err := os.Stat("file.txt")

	if err != nil {
		return
	}

	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size())
	fmt.Println("IsDir:", info.IsDir())
	fmt.Println("Permissions:", info.Mode())
}
