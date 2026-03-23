package main

import "fmt"

func main () {
	const age = 30

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(age, port, host);
}
