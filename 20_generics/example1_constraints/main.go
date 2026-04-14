package main

import "fmt"

type Number interface {
	int | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(10, 20))
	fmt.Println(Add(10.5, 2.3))
}
