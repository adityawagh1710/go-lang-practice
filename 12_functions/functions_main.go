package main

import (
	"errors"
	"fmt"
)

// Normal function with return type
func fun() int {
	return 123
}

// Functions with similar inputs Immediately Invoked
func addition(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

// Functions with multiple outputs
func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// Named returns: returns min and max
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return
}

func main() {

	fmt.Println(fun())

	fmt.Println(addition(1, 2))

	result, err := divide(10, 2)
	// Use err if its nil or not
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	_, err2 := divide(5, 0)
	// Use _ to ignore the result
	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	// Use array for min max
	lo, hi := minMax([]int{3, 1, 7, 2, 9, 4})
	fmt.Println("minMax result Min:", lo, "Max:", hi)

	// Anonymous functions have NO name
	square := func(n int) int {
		return n * n
	}
	fmt.Println("Anonymous functions output ", square(5))
	fmt.Println("Anonymous functions output ", square(9))

	// Store functions using map
	operations := map[string]func(int, int) int{
		"add":      addition,
		"subtract": subtract,
		"multiply": multiply,
	}

	fmt.Println(operations["add"](10, 5))
	fmt.Println(operations["subtract"](10, 5))
	fmt.Println(operations["multiply"](10, 5))
}
