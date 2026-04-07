package main

import "fmt"

type emp struct {
	Name string
	Age  int
}

func NewEmp(name string, age int) (*emp, error) {

	if age < 0 {
		return nil, fmt.Errorf("invalid age")
	}

	return &emp{Name: name, Age: age}, nil
}
