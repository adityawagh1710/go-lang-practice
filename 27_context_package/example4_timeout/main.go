package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Data fetched")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err := fetchData(ctx)

	fmt.Println("Result:", err)
}
