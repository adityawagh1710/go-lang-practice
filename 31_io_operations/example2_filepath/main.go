package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create file path safely (cross-platform)
	path := filepath.Join("data", "test.txt")

	// Create directory if not exists
	os.MkdirAll("data", os.ModePerm)

	// Create file
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Write to file
	file.WriteString("Hello Go File!\n")

	// Read file
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Delete file
	// err = os.Remove(path)

	// if err != nil {
	// 	panic(err)
	// }
}
