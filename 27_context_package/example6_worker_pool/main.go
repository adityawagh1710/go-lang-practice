package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case job := <-jobs:
			fmt.Println("Processing:", job)
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("Worker exiting")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan int)

	go worker(ctx, jobs)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	cancel()

	time.Sleep(1 * time.Second)
}
