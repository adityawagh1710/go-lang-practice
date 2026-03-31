package main

import (
	"fmt"
	"sync"
)

// worker simulates processing a job — squares the value
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		result := j * j
		fmt.Printf("  Worker %d processed job %d --->>> %d\n", id, j, result)
		results <- result
	}
}

func main() {
	jobs := make(chan int, 10) // buffered: producer won't block immediately
	results := make(chan int, 10)

	var wg sync.WaitGroup

	// Fan-Out: launch 3 worker goroutines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send 9 jobs into the jobs channel
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // signal workers: no more jobs coming

	// Wait for all workers, then close results channel
	go func() {
		wg.Wait()
		close(results) // safe to close here — all senders are done
	}()

	// Fan-In: collect all results
	total := 0
	for r := range results {
		total += r
	}
	fmt.Println("Total:", total) // 1+4+9+16+25+36+49+64+81 = 285
}
