package singleresponsiblity

import (
	"fmt"
	"time"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type UserRegister struct{}

func (ur *UserRegister) RegisterUser(user *User) {
	ur.saveUserToDB(user)
	ur.sendWelcomeEmail(user)
}

func (ur *UserRegister) saveUserToDB(user *User) {
	fmt.Printf("Saving user %s to the database...\n", user.FirstName)
	time.Sleep(1 * time.Second)
	fmt.Printf("User %s saved to the database successfully \n", user.FirstName)
}

func (ur *UserRegister) sendWelcomeEmail(user *User) {
	fmt.Printf("Sending welcome email to %s at %s \n", user.FirstName, user.Email)
	time.Sleep(1 * time.Second)
	fmt.Printf("Welcome email to %s at %s \n", user.FirstName, user.Email)
}

func NoSingleResponsiblityGo() {
	user := &User{
		ID:        1,
		FirstName: "Arpan",
		LastName:  "Das",
		Email:     "dasarpan471@gmail.com",
	}
	ur := &UserRegister{}
	ur.RegisterUser(user)
}

// The UserRegister struct is responsible for handling both user registration
// sending a welcome email
// this violates the Single Responsibility Principle because the UserRegister class has multiple responsibilities.
