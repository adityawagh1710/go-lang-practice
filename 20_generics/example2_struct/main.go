package main

import "fmt"

type Box[T any] struct {
	Value T
}

func main() {
	b1 := Box[int]{Value: 100}
	b2 := Box[string]{Value: "Generics"}

	fmt.Println(b1.Value)
	fmt.Println(b2.Value)
}
