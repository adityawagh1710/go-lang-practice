package main

import (
	"fmt"
	"time"
)

// Run using command go run *.go

func main() {
	// Declare a variable of type example set to its
	// zero value.
	var ex1 example
	// Display the value.
	fmt.Printf("%+v\n", ex1)

	// Declare a variable of type example and init using
	// a struct literal.
	ex2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	// Display the field values.
	fmt.Println("Flag", ex2.flag)
	fmt.Println("Counter", ex2.counter)
	fmt.Println("Pi", ex2.pi)

	var ex3 example
	// ex3 = ex2 Not allowed, compiler error
	// Allowed, NO compiler error
	ex3 = example(ex2)
	fmt.Printf("%+v\n", ex3)

	var ext2 struct {
		flag    bool
		counter int16
		pi      float32
	}
	// Allowed, NO need for conversion syntax
	ex1 = ext2
	fmt.Printf("%+v\n", ex1)

	p1 := person{"Rahul", "Kumar"}
	p1.details()
	e1 := employee{person: person{"John", "Ponting"}, empId: 11}
	e1.details()

	u := User{
		BaseModel: BaseModel{
			Id: 1,
			Timestamps: Timestamps{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		Username: "alice",
		Email:    "alice@go.dev",
	}

	fmt.Println(u.Id) // 1 — from BaseModel
	fmt.Println(u.CreatedAt)
	fmt.Println(u.Username)
	fmt.Println(u.Email)
	fmt.Println(u.UpdatedAt) // ... — from Timestamps (1 levels deep)
	u.Touch()                // calls Timestamps.Touch via promotion
	fmt.Println(u)           // ... — from Timestamps (2 levels deep)
}
