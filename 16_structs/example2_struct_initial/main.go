package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func (e Employee) SayHello() {
	fmt.Println("Hello, my name is", e.FirstName)
}

func main() {
	var x Employee
	x.FirstName = "TESTST"
	x.Age = 22
	fmt.Println(x)

	var john = Employee{"John", "Doe", 20}
	john.SayHello()
	fmt.Println(john)

	fmt.Println()

	var jane = &Employee{"Jane", "Doe", 22}
	jane.SayHello()

	// The data, labeled as a pointer.
	fmt.Println(jane)
	// The actual memory address (hexadecimal)
	fmt.Printf("%p", jane)
	// The data with field names included.
	fmt.Printf("%+v", jane)
	fmt.Println()

	var jane2 *Employee
	fmt.Println(jane2)
}
