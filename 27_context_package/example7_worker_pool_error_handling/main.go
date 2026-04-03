package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
	Err   error
}

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: shutting down\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return // Channel closed
			}
			// Simulate work with possible error
			time.Sleep(300 * time.Millisecond)
			if job.ID%3 == 0 {
				results <- Result{JobID: job.ID, Err: fmt.Errorf("failed job %d", job.ID)}
				continue
			}
			results <- Result{JobID: job.ID, Value: job.ID * 10}
		}
	}
}

func main() {
	// Step 1 : Create context package for timeouts
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Step 2 : Create channel for goroutine workers
	jobs := make(chan Job, 20)
	results := make(chan Result, 20)

	// Step 3 : Create sync WaitGroup for multiple goroutines
	var wg sync.WaitGroup

	// Start fixed workers (limits concurrency)
	const numWorkers = 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// Producer: send jobs
	for i := 1; i <= 15; i++ {
		select {
		case <-ctx.Done():
			goto shutdown
		default:
			jobs <- Job{ID: i}
		}
	}
	close(jobs) // Signal no more jobs

	// Collect results in background (fan-in)
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consume results
	for res := range results {
		if res.Err != nil {
			fmt.Printf("Error: %v\n", res.Err)
		} else {
			fmt.Printf("Result from job %d: %d\n", res.JobID, res.Value)
		}
	}

shutdown:
	fmt.Println("All done or cancelled.")
}
