package main

import (
	"context"
	"fmt"
)

type key string

func main() {
	ctx := context.WithValue(context.Background(), key("userID"), 42)

	process(ctx)
}

func process(ctx context.Context) {
	userID := ctx.Value(key("userID"))

	fmt.Println("User:", userID)
}
