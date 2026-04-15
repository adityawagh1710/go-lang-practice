package main

import (
	"fmt"
)

func Apply(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func Multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Map(nums []int, fn func(int) int) []int {
	res := make([]int, len(nums))

	for i, n := range nums {
		res[i] = fn(n)
	}

	return res
}

func main() {
	// higher-order functions
	fmt.Println("Apply1", Apply(2, 3, add))
	fmt.Println("Apply2", Apply(2, 3, mul))

	fmt.Println()

	// closures
	double := Multiplier(2)
	triple := Multiplier(3)

	fmt.Println("Multiplier1", double(5)) // 10
	fmt.Println("Multiplier2", triple(5)) // 15

	fmt.Println()

	// map function
	nums := []int{1, 2, 3}

	squared := Map(nums, func(n int) int {
		return n * n
	})

	fmt.Println("squared:", squared)

	cubed := Map(nums, func(n int) int {
		return n * n * n
	})

	fmt.Println("cubed:", cubed)

	fmt.Println()
}
