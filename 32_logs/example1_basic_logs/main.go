package main

import (
	"log"
	"os"
)

func main() {
	// Default logger writes to stderr
	log.Println("This is a normal log")
	log.Print("Another log")

	log.SetPrefix("APP: ")
	log.Println("With prefix")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Application started")

	log.Printf("User %s logged in at port %d", "aditya", 8080)

	// Custom logger with prefix and flags
	infoLogger := log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile)

	errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)

	infoLogger.Println("Server is running")
	errorLogger.Println("Database connection failed")

	// Fatal calls os.Exit(1) after logging
	// log.Fatal("unrecoverable error")

	// Panic calls panic() after logging
	// log.Panic("panic situation")

}
