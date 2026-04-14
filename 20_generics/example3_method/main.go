package main

import "fmt"

type Box[T any] struct {
	Value T
}

func (b Box[T]) Get() T {
	return b.Value
}

func main() {
	b := Box[int]{Value: 42}

	fmt.Printf("%#v\n", b.Get())

	b1 := Box[string]{Value: "100"} // This will cause a compile-time error

	fmt.Printf("%#v\n", b1.Get())
}
