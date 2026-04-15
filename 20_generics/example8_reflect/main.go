package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func Add(a, b int) int {
	return a + b
}

type User2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// basic reflection
	x := 10
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t) // int
	fmt.Println(v) // 10

	fmt.Println()

	// working with pointers
	var x1 int32 = 10
	t1 := reflect.TypeOf(x1)
	fmt.Println(t1)        // int32
	fmt.Println(t1.Kind()) // int

	fmt.Println()

	// iterating over struct fields using reflection
	u := User{"Adi", 25}
	v1 := reflect.ValueOf(u)
	for i := 0; i < v1.NumField(); i++ {
		fmt.Println(v1.Field(i))
	}

	fmt.Println()

	// calling function using reflection
	fn := reflect.ValueOf(Add)
	args := []reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf(3),
	}
	result := fn.Call(args)
	fmt.Println("Reflect result:", result[0].Int()) // 5

	fmt.Println()

	// struct tags
	t2 := reflect.TypeOf(User2{})
	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		fmt.Println(field.Tag.Get("json"))
	}

	fmt.Println()
}
