package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Channels")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// this goroutine is for reading
	// Read only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	// this goroutine is for writing
	// Write only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5

		close(ch)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
