package main

import "fmt"

// enumerated types
type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared  OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)

	fmt.Println("CONST 1:", A1, B1, C1)

	fmt.Println("CONST 2:", A2, B2, C2)

	fmt.Println("CONST 3:", A3, B3, C3)

	fmt.Println("LOGS :", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)

	logger := &Logger{Level: Info} // Only show Info and above
	logger.Log(Debug, "This debug message will be hidden")
	logger.Log(Info, "Application started successfully")
	logger.Log(Warning, "Disk space is running low")
	logger.Log(Error, "Failed to connect to database")
	logger.Log(Fatal, "Critical failure — shutting down")

	// Type safety in action:
	// logger.Level = 99          // Compile error! (cannot use int as LogLevel)
	// logger.Level = Debug       // OK
	// Printing the enum directly
	fmt.Println("Current level:", Info)         // Output: Current level: INFO
	fmt.Printf("Level as int: %d\n", int(Info)) // 1
}
