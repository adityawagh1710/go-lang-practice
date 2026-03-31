package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count int
var mutex sync.Mutex

func increment(s string) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "Count: ", count)
		mutex.Unlock()

	}
	wait.Done()

}

var counter2 int

func increment2(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter2++
		mutex.Unlock()
	}
}

func main() {
	fmt.Println("Basic examples 1")
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)

	fmt.Println("Basic examples 2")

	for i := 0; i < 2; i++ {
		wait.Add(1)
		go increment2(&wait)
	}

	wait.Wait()
	fmt.Println(counter2) // correct result
}
