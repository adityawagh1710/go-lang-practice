package main

import "fmt"

const (
	A1 = iota // 0 : Start at 0
	B1 = iota // 1 : Increment by 1
	C1 = iota // 2 : Increment by 1
)

const (
	A2 = iota // 0 : Start at 0
	B2        // 1 : Increment by 1
	C2        // 2 : Increment by 1
)

const (
	A3 = iota + 1 // 1 : Start at 0 + 1
	B3            // 2 : Increment by 1
	C3            // 3 : Increment by 1
)

const (
	Ldate         = 1 << iota //  1 : Shift 1 to the left 0.  0000 0001
	Ltime                     //  2 : Shift 1 to the left 1.  0000 0010
	Lmicroseconds             //  4 : Shift 1 to the left 2.  0000 0100
	Llongfile                 //  8 : Shift 1 to the left 3.  0000 1000
	Lshortfile                // 16 : Shift 1 to the left 4.  0001 0000
	LUTC                      // 32 : Shift 1 to the left 5.  0010 0000
)

type LogLevel int

// Using iota to define the enum values.
// Starting from 0 is common, but you can do iota + 1 if you prefer 1-based.
const (
	Debug   LogLevel = iota // 0
	Info                    // 1
	Warning                 // 2
	Error                   // 3
	Fatal                   // 4
)

// String() implements the fmt.Stringer interface.
// This is what makes fmt.Println(level) print "INFO" instead of "1".
// Very important for logging, debugging, and JSON output.
func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Optional: Add a method for colored console output (common in real loggers).
func (l LogLevel) Color() string {
	switch l {
	case Debug:
		return "\033[36m" // Cyan
	case Info:
		return "\033[32m" // Green
	case Warning:
		return "\033[33m" // Yellow
	case Error:
		return "\033[31m" // Red
	case Fatal:
		return "\033[35m" // Magenta
	default:
		return "\033[0m"
	}
}

// Example of a simple logger using our enum.
type Logger struct {
	Level LogLevel
}

func (l *Logger) Log(level LogLevel, msg string) {
	if level < l.Level {
		return // filter lower levels
	}
	fmt.Printf("%s[%s]\033[0m %s\n", level.Color(), level, msg)
}
