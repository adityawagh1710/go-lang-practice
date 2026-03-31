package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0

	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("Circle area: %.2f\n", area)
}
