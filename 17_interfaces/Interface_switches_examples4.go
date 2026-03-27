package main

import "fmt"

func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
	case int:
		fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
	default:
		fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
	}
}

