package main

import (
	"fmt"
	"time"
)

type person struct {
	fname string
	lname string
}

// Struct is a data type and can be used as an anonymous field (having only the type).
// One struct can be inserted or "embedded" into other struct.
// People is in employee
type employee struct {
	person
	empId int
}

func (p person) details() {
	fmt.Println(p, " "+" I am a person")
}

func (e employee) details() {
	fmt.Println(e, " "+"I am a employee")
}

type example struct {
	flag    bool
	counter int16
	pi      float32
}

type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Timestamps) Touch() {
	t.UpdatedAt = time.Now()
}

type BaseModel struct {
	Id         int
	Timestamps // embedded
}

type User struct {
	BaseModel // embedded — ID, CreatedAt, UpdatedAt all promoted
	Username  string
	Email     string
}
