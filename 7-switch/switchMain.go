package main 

import "fmt"
import "time"

// Type switch 
func process(data interface{}) {
    switch v := data.(type) {
    case string:
        fmt.Println("Length:", len(v))
    case int:
        fmt.Println("Double:", v*2)
    case []int:
        fmt.Println("Slice length:", len(v))
    default:
        fmt.Println("Unknown type")
    }
}

func main ()  {
	i := 3;

	switch i {
	case 1:
		fmt.Println("kjaSkdj---1")
	case 2:
		fmt.Println("kjaSkdj---2")
	case 3:
		fmt.Println("kjaSkdj---3")
	case 4:
		fmt.Println("kjaSkdj---4")
	}


	// Multi condition switch 
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("week")
	}

	process("hello")
    process(10)
    process([]int{1, 2, 3, 4})
	 // default case
    process(true)
}