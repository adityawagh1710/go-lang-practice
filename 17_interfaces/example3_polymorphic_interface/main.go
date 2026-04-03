package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Printf("%s: Woof!\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Printf("%s: Meow!\n", c.Name)
}

// MakeNoise accepts any []Speaker — decoupled from concrete types
func MakeNoise(speakers []Speaker) {
	for _, s := range speakers {
		s.Speak()
	}
}

// NotifyAll accepts a single Speaker — also decoupled
func NotifyAll(s Speaker) {
	s.Speak()
}

func main() {
	group := []Speaker{
		Dog{Name: "Rex"},
		Cat{Name: "Luna"},
	}

	MakeNoise(group)
	// MakeNoise works for ANY type with Speak() — no code changes needed
	// to add a new animal, robot, or any other type in the future
}
