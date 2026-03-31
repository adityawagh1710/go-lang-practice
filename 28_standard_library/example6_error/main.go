package main

import (
	"errors"
	"fmt"
)

func divideNumbers(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divideNumbers(10, 1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
