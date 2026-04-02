package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second): // simulates slow operation
		fmt.Println("fetch: data ready")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {

	deadline := time.Now().Add(1 * time.Second) // deadline 1s from now

	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	err := fetchData(ctx)

	if err != nil {
		fmt.Println("error:", err) // context deadline exceeded
	}

	if dl, ok := ctx.Deadline(); ok {
		fmt.Println("deadline was:", dl.Format(time.RFC3339))
	}
}
