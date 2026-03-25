package main

import "fmt"

// Basic variadics
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// Zero arguments
func printNumbers(nums ...int) {
	fmt.Println(nums)
}

func display(nums ...int) {
    fmt.Println("Slice:", nums)
    fmt.Println("Length:", len(nums))
}

func main() {
	// Basic variadics
	fmt.Println(sum(1, 2, 3, 6, 9, 10))

	fmt.Println(sum(10, 20))

	// Zero arguments
	printNumbers()

	// Variadics with slice
	numbers := []int{1, 2, 3}
	fmt.Println(sum(numbers...))

	display(numbers...)
}
