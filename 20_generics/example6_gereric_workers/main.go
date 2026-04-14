package main

import (
	"fmt"
	"sync"
)

type Result[T any] struct {
	Value T
}

func worker[T any](wg *sync.WaitGroup, input T, out chan Result[T]) {
	defer wg.Done()
	out <- Result[T]{Value: input}
}

func main() {
	var wg sync.WaitGroup
	
	ch := make(chan Result[int], 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(&wg, i, ch)
	}

	wg.Wait()
	close(ch)

	for res := range ch {
		fmt.Println(res.Value)
	}
}
