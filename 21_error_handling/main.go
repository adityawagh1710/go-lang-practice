package main

import (
	"errors"
	"fmt"
)

var (
	// ErrBadRequest is returned when there are problems with the request.
	ErrBadRequest = errors.New("Bad Request")

	// ErrPageMoved is returned when a 301/302 is returned.
	ErrPageMoved = errors.New("Page Moved")
)

func main() {

	err := errors.New("an error occurred")
	fmt.Println(err)

	if err := webCall(false); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("Bad Request Occurred")
			return

		case ErrPageMoved:
			fmt.Println("The Page moved")
			return

		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Life is good")

}

// webCall performs a web operation.
func webCall(b bool) error {
	if b {
		return ErrBadRequest
	}

	return ErrPageMoved
}
