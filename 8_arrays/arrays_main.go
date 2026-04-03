package main

import "fmt"

func main() {
	var nums [4]int

	// 0th index as 12
	nums[0] = 12

	// Array length
	fmt.Println("length ", len(nums))

	// All array element
	fmt.Println(nums)

	// Declare it in single line
	nums = [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	// 2D array (use different variable)
	matrix := [2][2]int{{3, 8}, {4, 4}}
	fmt.Println(matrix)

	// Fixed size, that is predictable
	// Memeory optimization
}
