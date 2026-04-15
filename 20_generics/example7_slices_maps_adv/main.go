package main

import (
	"fmt"
	"slices"
)

type User struct {
	Name string
	Age  int
}

func main() {
	nums := []int{5, 2, 8, 1, 3}
	slices.Sort(nums)
	fmt.Println("sorted nums:", nums) // [1 2 3 5 8]

	users := []User{
		{"A", 30},
		{"B", 20},
		{"C", 25},
	}

	slices.SortFunc(users, func(a, b User) int {
		return a.Age - b.Age
	})

	fmt.Println("sorted users:", users)

	arr := []int{1, 2, 3}
	fmt.Println("slices.Contains(arr, 2):", slices.Contains(arr, 2)) // true
	fmt.Println("slices.Contains(arr, 4):", slices.Contains(arr, 4)) // false

	arr2 := []string{"go", "java", "python"}
	idx := slices.Index(arr2, "java")
	fmt.Println("idx:", idx)

}
