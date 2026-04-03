package main

import "fmt"

func main() {
	age := 16

	// Simple if, else if
	if age > 30 {
		fmt.Println("adult")
	} else if age > 12 {
		fmt.Println("Teen")
	} else {
		fmt.Println("child")
	}

	if age := 20; age >= 18 {
		fmt.Println("Person adult ", age)
	} else if age >= 12 {
		fmt.Println("Person teenager ", age)
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" && hasPermissions {
		fmt.Println("ADMIN HAI")
	}

	// Go dont have ternary operators we need to use if else only
}
