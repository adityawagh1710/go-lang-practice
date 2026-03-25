package main

import "fmt"

func addMarks(m map[string]int) {
	m["Peter"] = 82
	m["Johnson"] = 94
}

func main() {
	// Method 1: Using make()
	m1 := make(map[string]int)
	fmt.Println(m1)

	a := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	fmt.Println(a)

	m := make(map[int]bool)
	m[1] = true
	m[2] = false
	m[3] = false
	fmt.Println(m)
	fmt.Println("The value:", m[3])

	// Method 2: Map literal
	// *** var a map[KeyType]ValueType ***
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m2)

	// Method 3: Retrieving Element with two values in Map
	m3 := make(map[string]int)
	m3["Key1"] = 15
	m3["Key2"] = 25
	m3["Key3"] = 35
	fmt.Println(m3)
	v, ok := m3["Key2"]
	fmt.Println("The value:", v, "Present?", ok)
	i, j := m3["Key4"]
	fmt.Println("The value:", i, "Present?", j)

	// Method 4: Maps as Function Parameters in Golang
	students := make(map[string]int)
	addMarks(students)
	fmt.Println(students)
}
