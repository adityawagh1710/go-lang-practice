package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func increment(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func changePtr(ptr *int) {
	// `*int` → pointer to int
	*ptr = 10
}

func main() {
	// "Pass by ref"
	x, y := 10, 20
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After : x=%d, y=%d\n", x, y)

	// Display the "value of" and "address of" count.
	// "Pass by value" of the count.
	count := 10
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	increment(count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// `*int` → pointer to int
	ptr := new(int)
	fmt.Println("Before change ptr", *ptr)
	changePtr(ptr)
	fmt.Println("After change ptr", *ptr)
}
