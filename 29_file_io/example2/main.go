package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// MultiWriter: write to multiple destinations simultaneously
	logFile, _ := os.Create("log.txt")

	defer logFile.Close()

	// Writes to both stdout and log file at the same time
	multi := io.MultiWriter(os.Stdout, logFile)

	fmt.Fprintln(multi, "This goes to console AND file simultaneously")

	// MultiReader: concatenate multiple readers into one
	r1 := strings.NewReader("Hello, ")

	r2 := strings.NewReader("World!\n")

	combined := io.MultiReader(r1, r2)

	data, _ := io.ReadAll(combined)

	fmt.Printf("Combined: %s", data)
}
