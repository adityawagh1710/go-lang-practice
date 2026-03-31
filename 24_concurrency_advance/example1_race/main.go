package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int

func increment(s string) {
	for i := 0; i < 10; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(4)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)

	}
	wg.Done()

}

func main() {
	wg.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wg.Wait()
	fmt.Println("last count value ", count)
}
