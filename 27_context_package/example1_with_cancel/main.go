package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("worker %d doing work...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // always defer — prevents goroutine leak

	go doWork(ctx, 1)
	go doWork(ctx, 2)

	time.Sleep(1500 * time.Millisecond)
	cancel() // signal all workers to stop

	time.Sleep(200 * time.Millisecond) // let workers print their exit message
	fmt.Println("main: done")
}
