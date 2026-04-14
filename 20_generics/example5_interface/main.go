package main

import "fmt"

type Printer interface {
	Print() string
}

type User struct {
	Name string
}

func (u User) Print() string {
	return u.Name
}

func PrintData[T Printer](p T) {
	fmt.Println(p.Print())
}

func main() {
	u := User{Name: "Aditya"}

	PrintData(u)
}
