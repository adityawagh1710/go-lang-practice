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

func main() {
	x, y := 10, 20

	fmt.Printf("Before: x=%d, y=%d\n", x, y)

	// "Pass by ref"
	swap(&x, &y)

	fmt.Printf("After:  x=%d, y=%d\n", x, y)

	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// "Pass by value" of the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}
