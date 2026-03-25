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

	fmt.Println("Range with defining the order and oop with no order")

	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	// defining the order
	var b []string
	b = append(b, "one", "two", "three", "four")

	// loop with no order
	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()
	// loop with the defined order
	// Ignore index with _,
	for _, element := range b {
		fmt.Printf("%v : %v, ", element, a[element])
	}

	fmt.Println()
}
