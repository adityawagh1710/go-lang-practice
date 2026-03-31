package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func() (string, error) // Any function that returns result + error

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("[Worker %d] executing task...\n", id)
		result, err := task()
		if err != nil {
			results <- fmt.Sprintf("[Worker %d] ERROR: %v", id, err)
			continue
		}
		results <- fmt.Sprintf("[Worker %d] SUCCESS: %s", id, result)
	}
}

func main() {
	const numWorkers = 4
	tasks := make(chan Task, 20)
	results := make(chan string, 20)
	var wg sync.WaitGroup

	// Launch workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Generate sample tasks (could be API calls, DB queries, etc.)
	for i := 1; i <= 15; i++ {
		taskID := i
		tasks <- func() (string, error) {
			time.Sleep(800 * time.Millisecond) // simulate work
			if taskID%5 == 0 {
				return "", fmt.Errorf("task %d failed", taskID)
			}
			return fmt.Sprintf("Task %d completed", taskID), nil
		}
	}

	close(tasks)   // No more tasks
	wg.Wait()      // Wait for workers
	close(results) // Close results channel

	// Collect results
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("Generic worker pool finished.")
}
