package main

import (
	"fmt"
	"slices"
)

func main() {
	// Slice with length = 3, capacity = 5
	nums := make([]int, 3, 5)

	fmt.Println("Slice length capacity ", nums)
	fmt.Println("len ", len(nums))
	fmt.Println("cap ", cap(nums))

	// Access and update array by index
	// slice_name := []datatype{values}
	nums2 := []int{10, 20, 30}
	fmt.Println("Slice 0 index ", nums2[0])
	nums2[1] = 50
	fmt.Println("Slice 1 index ", nums2)

	// Slice append to attach multiple array elements
	// slice_name = append(slice_name, element1, element2, ...)
	nums3 := make([]int, 0, 5)
	nums3 = append(nums3, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println("Append ", nums3)

	// Append array elemnts for slice
	nums4 := make([]int, 0, 5)
	nums4 = append(nums4, len(nums4))
	nums4 = append(nums4, 7)
	fmt.Println("Append slice", nums4)

	// Copy slice from one to new
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src)
	fmt.Println("Copy ", src, dest)

	// Slice operator
	nums5 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Slice 0-2", nums5[0:2])
	fmt.Println("Slice all after 2", nums5[:2])
	fmt.Println("Slice only first 2", nums5[2:])

	// Slice package
	nums6 := []int{1, 2, 3, 4, 5, 7}
	nums7 := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("Slice equal", slices.Equal(nums6, nums7))
}
