package main

import "fmt"

func main() {
	fmt.Println("Welcome to myStruct")
	arpan := User{"Arpan Das", "arpandas@go.dev", true, 16}

	fmt.Println(arpan)
	fmt.Printf("hitesh details are: %+v\n", arpan)
	fmt.Printf("Name is %v and email is %v.", arpan.Name, arpan.Email)
}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}
