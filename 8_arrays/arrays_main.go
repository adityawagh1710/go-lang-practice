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

	// declares two arrays (arr1 and arr2) with inferred lengths
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// initializes only the second and third elements of the array
	arr3 := [5]int{1: 10, 2: 40}
	fmt.Println(arr3)
}
