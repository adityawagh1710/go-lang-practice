package main

import (
	"fmt"
	"time"
)

func main() {
	// Current time
	now := time.Now()

	fmt.Println("Now:", now)

	// Format time
	formatted := now.Format("2006-01-02 15:04:05")

	fmt.Println("Formatted:", formatted)

	// Parse time
	str := "2026-04-02 18:30:00"
	
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Parsed:", parsedTime)

	// Duration
	start := time.Now()

	time.Sleep(2 * time.Second)
	
	elapsed := time.Since(start)

	fmt.Println("Elapsed:", elapsed)
}
