package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func SendPerson(ch chan Person, p Person) {
	ch <- p
}

func main() {

	// waitForResult: In this pattern, the parent goroutine waits for the child
	// goroutine to finish some work to signal the result with struct
	personStruct := Person{"John", 23}

	ch := make(chan Person)

	go SendPerson(ch, personStruct)

	val, ok := <-ch

	fmt.Println(val, ok)

	if !ok {
		fmt.Println("Channel closed")
	}

	name := val.Name

	fmt.Println(name)
}
