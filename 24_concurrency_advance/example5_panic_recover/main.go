package main

import "fmt"

func riskyOperation() {
	defer fmt.Println("Cleanup: resources released")
	panic("Unexpected error in riskyOperation")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Main recovered:", r)
		}
	}()

	fmt.Println("Starting program...")
	riskyOperation()
	fmt.Println("Program completed successfully")
}

// Output:
// Starting program...
// Cleanup: resources released
// Main recovered: Unexpected error in riskyOperation
