package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// Define flags: flag.Type("name", defaultValue, "description")
	port := flag.Int("port", 8080, "HTTP server port")
	debug := flag.Bool("debug", false, "Enable debug mode")
	name := flag.String("name", "World", "Name to greet")
	timeout := flag.Duration("timeout", 30*time.Second, "Request timeout")

	flag.Parse() // Must call Parse() before reading values

	fmt.Printf("Port: %d\n", *port) // dereference pointer
	fmt.Printf("Debug: %v\n", *debug)
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Timeout: %v\n", *timeout)

	// Remaining non-flag args (e.g. filenames after flags)
	args := flag.Args()
	
	fmt.Println("Extra args:", args)
}
