package main

import "fmt"

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Printf("Printer Name: %s\n", e.name)
}
