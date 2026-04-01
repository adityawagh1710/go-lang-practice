package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read and write to another file (streaming fashion)

	logFile, _ := os.Create("example.txt")

	defer logFile.Close()

	logFile.WriteString("Hi this is go long Yahoo")

	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)

	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		e := writer.WriteByte(b)

		if e != nil {
			panic(e)
		}

	}

	writer.Flush()

	fmt.Println("written to new file succesfully")

	err = os.Remove("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("example deleted")
}
