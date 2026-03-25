package main

import "fmt"

func main() {
	// Iterating over map and slices
	fmt.Println("Range map")

	m := map[string]string{"a": "a", "b": "b", "c": "c"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Range unicode code point rune")
	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
	for k, v := range "johndove" {
		fmt.Println(k, string(v))
	}

	fmt.Println("Range slices")

	nums := []int{6, 7, 8}
	for i, num := range nums {
		fmt.Println(num, i)
	}
}
