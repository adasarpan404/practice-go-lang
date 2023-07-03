package main

import "fmt"

func main() {
	fmt.Println("Welcome to myStruct")
	arpan := User{"Arpan Das", "arpandas@go.dev", true, 16}

	fmt.Println(arpan)

}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}
