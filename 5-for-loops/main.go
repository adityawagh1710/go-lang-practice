package main
import "fmt"

func main ()  {
	i := 1
	
	// Normal loop
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Infinite loop
	//	for {
	//		fmt.Println(101010);
	//	}

	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}