package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker:", id)
}

func main() {
	fmt.Println("Basic goroutine")

	time.Sleep(1000 * time.Millisecond)
	go sayHello("Aditya")

	time.Sleep(1000 * time.Millisecond)
	go sayHello("Go World")

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Basic goroutine done")

	fmt.Println("Basic WaitGroup example 1")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Basic WaitGroup example 2")

	// wg is used to manage concurrency.
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Create a goroutine from the lowercase function.
	go func() {
		fmt.Println("Waiting anonymous function 1")
		lowercase()
		wg.Done()
	}()

	// Create a goroutine from the uppercase function.
	go func() {
		fmt.Println("Waiting anonymous function 2")
		uppercase()
		wg.Done()
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println()
	fmt.Println("DONE")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}

	fmt.Println("")
}

// uppercase displays the set of uppercase letters three times.
func uppercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}

	fmt.Println("")
}
