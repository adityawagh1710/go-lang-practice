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

	fmt.Println(now.Format("2006-01-02"))              // 2026-04-01
	fmt.Println(now.Format("02-Jan-2006"))             // 01-Apr-2026
	fmt.Println(now.Format("15:04:05"))                // 20:33:09
	fmt.Println(now.Format("Monday, January 2, 2006")) // Wednesday, April 1, 2026
	fmt.Println(now.Format(time.RFC3339))              // 2026-04-01T20:33:09+05:30
}
