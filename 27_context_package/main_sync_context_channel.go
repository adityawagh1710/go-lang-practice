package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done() // (A)

	for {
		select {
		case job, ok := <-jobs: // (B)
			if !ok {
				fmt.Println("Worker", id, "done")
				return
			}

			time.Sleep(time.Microsecond) // simulate work

			mu.Lock() // (C)

			counter++

			fmt.Printf("(Worker %d) (processed job %d) (counter=%d) (at %v) \n", id, job, counter, time.Now().Format("2006-01-02 15:04:05"))

			mu.Unlock()

		case <-ctx.Done(): // (D)
			fmt.Println("Worker", id, "stopped:", ctx.Err())
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // (E)

	defer cancel()

	jobs := make(chan int) // (F)

	var wg sync.WaitGroup // (G)

	// start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)                    // (H)
		go worker(ctx, i, jobs, &wg) // (I)
	}

	// send jobs
	go func() {
		for i := 1; i <= 1000; i++ {
			jobs <- i // (J)
		}
		close(jobs)
	}()

	wg.Wait() // (K)

	fmt.Println("Final Counter:", counter)
}
