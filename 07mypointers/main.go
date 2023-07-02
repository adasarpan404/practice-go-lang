package main

import "fmt"

func main() {
	fmt.Println("Welcome to my pointer")

	// var ptr *int
	// fmt.Println("the value of ptr =", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("the actual value of ptr =", ptr)
	fmt.Println("the actual value of ptr =", *ptr)

	*ptr = *ptr * 2

	fmt.Println("the actual value of ptr =", *ptr)
}
