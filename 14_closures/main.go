package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

func addFixed(x int) int {
	return x + 10
}

func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x // 'base' is captured
	}
}

func main() {
	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	// Independent state
	counter2 := counter()
	fmt.Println(counter2())

	add10 := makeAdder(10)
	add20 := makeAdder(20)
	fmt.Println(add10(5))
	fmt.Println(add20(5))
	fmt.Println(addFixed(5))
}

/*
	Go supports anonymous functions, also known as closures.
	These are functions without a name, and they can be assigned to variables and used as arguments to other functions.
	Closures are often used in Go for tasks like defining inline functions or for concurrent programming using goroutines.
*/
