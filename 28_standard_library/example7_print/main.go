package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	name := "Aditya"
	age := 25

	// --------------------------------------------------
	// PRINT FAMILY (direct output)
	// --------------------------------------------------

	// Print (no newline, no formatting)
	fmt.Print("Print: ")
	fmt.Print(name, " ", age)
	fmt.Print("\n")

	// Println (adds newline, spaces auto)
	fmt.Println("Println:", name, age)

	// Printf (formatted output)
	fmt.Printf("Printf: Name=%s Age=%d\n", name, age)

	// --------------------------------------------------
	// STRING FAMILY (returns string)
	// --------------------------------------------------

	// Sprint (no formatting)
	s1 := fmt.Sprint("Sprint:", name, age)
	fmt.Println(s1)

	// Sprintln (adds newline in string)
	s2 := fmt.Sprintln("Sprintln:", name, age)
	fmt.Print(s2) // already has newline

	// Sprintf (formatted string)
	s3 := fmt.Sprintf("Sprintf: Name=%s Age=%d", name, age)
	fmt.Println(s3)

	// --------------------------------------------------
	// WRITER FAMILY (writes to io.Writer)
	// --------------------------------------------------

	// 1) Writing to STDOUT
	fmt.Fprint(os.Stdout, "Fprint to Stdout: ", name, "\n")

	fmt.Fprintln(os.Stdout, "Fprintln to Stdout:", name, age)

	fmt.Fprintf(os.Stdout, "Fprintf to Stdout: Name=%s Age=%d\n", name, age)

	// 2) Writing to buffer (in-memory)
	var buf bytes.Buffer

	fmt.Fprint(&buf, "Fprint buffer: ", name, "\n")
	fmt.Fprintln(&buf, "Fprintln buffer:", name, age)
	fmt.Fprintf(&buf, "Fprintf buffer: Name=%s Age=%d\n", name, age)

	fmt.Println("\n--- Buffer Output ---")
	fmt.Println(buf.String())

	// 3) Writing to file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprint(file, "Fprint file: ", name, "\n")
	fmt.Fprintln(file, "Fprintln file:", name, age)
	fmt.Fprintf(file, "Fprintf file: Name=%s Age=%d\n", name, age)

	fmt.Println("\nData written to output.txt")

}
