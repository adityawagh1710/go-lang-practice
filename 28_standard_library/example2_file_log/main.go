package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Write data to file
	file.WriteString("Hello, Go lang")
}
