package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// waitForResult: In this pattern, the parent goroutine waits for the child
	// goroutine to finish some work to signal the result.

	// Example 1
	fmt.Println("Basic channels example 1")

	ch := make(chan int) // unbuffered

	go func() {
		time.Sleep(time.Duration(time.Second))

		ch <- rand.Intn(100)

		fmt.Println("example 1 goroutine sent signal")
	}()

	value := <-ch

	fmt.Println(value)

	// Example 2
	fmt.Println("Basic channels example 2")

	messages := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		messages <- "ping"

		fmt.Println("example 2 goroutine sent signal")
	}()

	msg := <-messages

	fmt.Println(msg)
}
