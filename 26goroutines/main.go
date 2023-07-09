package main

import "fmt"

func runner1() {
	fmt.Println("It is the first runner")
}
func runner2() {
	fmt.Println("It is the second runner")
}

func execute() {
	runner1()
	runner2()
}

func main() {
	execute()
}
