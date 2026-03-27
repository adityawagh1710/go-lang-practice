package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func main() {
	fmt.Println("==========NORMAL INTERFACE==========")
	var s Speaker = Dog{}
	fmt.Println(s.Speak())

}
