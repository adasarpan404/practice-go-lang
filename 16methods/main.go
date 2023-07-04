package main

import "fmt"

func main() {
	fmt.Println("Welcome to Method")
	arpan := User{"Arpan", "Email", true, 19}

	fmt.Println(arpan)
	fmt.Printf("hitesh details are: %+v\n", arpan)
	fmt.Printf("Name is %v and email is %v.\n", arpan.Name, arpan.Email)
	arpan.getStatus()
	arpan.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", arpan.Name, arpan.Email)

}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}

func (user User) getStatus() {
	fmt.Println("Is user active: ", user.status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
