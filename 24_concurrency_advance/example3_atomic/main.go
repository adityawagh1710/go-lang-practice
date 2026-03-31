package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

var count int64

func increment(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration((rand.Intn(3))) * time.Millisecond)

		atomic.AddInt64(&count, 1)

		fmt.Println(s, i, "Count ->", count)
	}
	wg.Done()
}

// counter is a variable incremented by all goroutines.
var counter int64

func main() {
	fmt.Println("=====================================================")

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	wg.Add(grs)

	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {
			for i := 0; i < 2; i++ {
				atomic.AddInt64(&counter, 1)
			}

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()

	// Display the final value.
	fmt.Println("Final Counter:", counter)

	fmt.Println("=====================================================")

	wg.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wg.Wait()
	fmt.Println("last count value ", count)

	fmt.Println("=====================================================")
}
