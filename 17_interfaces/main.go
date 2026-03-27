package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return " Dogggy says Woof"
}

// retrieveFile can read from a file and process the data.
func retrieveFile(f file) error {
	data := make([]byte, 100)

	len, err := f.read(data)

	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))

	return nil
}

// retrievePipe can read from a pipe and process the data.
func retrievePipe(p pipe) error {
	data := make([]byte, 100)

	len, err := p.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))

	return nil
}

func main() {
	fmt.Println("==========NORMAL INTERFACE============")

	var s Speaker = Dog{}

	fmt.Println(s.Speak())

	fmt.Println("==========Slice of INTERFACE==========")

	c := canon{"PIXMA TR4520"}
	e := epson{"WorkForce Pro WF-3720"}

	printers := []printer{
		c,
		&e,
	}

	c.name = "PROGRAF PRO-1000"
	e.name = "Home XP-4100"

	for _, p := range printers {
		p.print()
	}

	fmt.Println("==========Example 2 INTERFACE==========")

	rect := rect{width: 3, height: 4}
	circle := circle{radius: 5}

	measure(rect)
	measure(circle)

	detectCircle(rect)
	detectCircle(circle)

	fmt.Println("==========Example 3 polymorphic INTERFACE==========")

	// Create two values one of type file and one of type pipe.
	f := file{"data.json"}
	p := pipe{"cfg_service"}

	// Call each retrieve function for each concrete type.
	retrieveFile(f)
	retrievePipe(p)

	fmt.Println("==========Example 4 switches INTERFACE==========")

	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)
}
