package main

import "fmt"

// Stage 1: generator — produces integers into a channel
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n // send each number downstream
		}
		close(out) // always close when done sending
	}()
	return out
}

// Stage 2: square — receives ints, sends their squares
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { // range over channel reads until it's closed
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Wire up the pipeline: generate → square → print
	nums := generate(2, 3, 4, 5, 6, 7, 8, 9)
	squares := square(nums)

	// Stage 3: consumer — reads final output
	for result := range squares {
		fmt.Println(result) // 4, 9, 16, 25
	}
}
