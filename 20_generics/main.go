package main

import "fmt"

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// T is a type parameter that is used like normal type inside the function
// any is a constraint on type i.e T has to implement "any" interface
func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}

	return r
}

func ZipToMap[K comparable, V any](keys []K, vals []V) map[K]V {
	if len(keys) != len(vals) {
		panic("keys and vals must have same length")
	}

	result := make(map[K]V, len(keys))

	for i, k := range keys {
		result[k] = vals[i]
	}

	return result
}

func main() {
	fmt.Println("==========Example 1============")
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))

	fmt.Println("==========Example 2============")
	vals := []bool{true, false, true}
	printSlice(vals, "john")

	fmt.Println("==========Example 3============")
	names := []string{"Alice", "Bob", "Charlie"}
	scores := []int{90, 85, 92}
	m := ZipToMap(names, scores)
	fmt.Println(m)

	fmt.Println("==========Example 4============")
}
