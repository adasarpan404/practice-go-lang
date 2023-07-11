package main

import (
	"fmt"
	"sync"
)

type State struct {
	value int
}

func useState(initialValue int) (chan<- int, <-chan int) {
	stateCh := make(chan int)
	getCh := make(chan int)

	state := State{value: initialValue}

	go func() {
		for {
			select {
			case newValue := <-stateCh:
				state.value = newValue
			case getCh <- state.value:
			}
		}
	}()
	return stateCh, getCh
}

func useEffect(fn func()) {
	go fn()
}
func main() {
	stateCh, getCh := useState(0)

	fmt.Println("Initial state:", <-getCh)
	stateCh <- 10
	fmt.Println("Updated state:", <-getCh)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	useEffect(func() {
		fmt.Println("Running side effect")
		fmt.Println("Side effect completed")
		defer wg.Done()
	})
	wg.Wait()
}
