package singleresponsiblity

import (
	"fmt"
	"time"
)

type UserDB struct{}

func (udb *UserDB) SaveUser(user *User) {
	fmt.Printf("Saving user %s to the database...\n", user.FirstName)
	time.Sleep(1 * time.Second)
	fmt.Printf("User %s saved to the database successfully \n", user.FirstName)
}

type EmailSender struct{}

func (es *EmailSender) SendWelcomeEmail(user *User) {
	fmt.Printf("Sending welcome email to %s at %s \n", user.FirstName, user.Email)
	time.Sleep(1 * time.Second)
	fmt.Printf("Welcome email to %s at %s \n", user.FirstName, user.Email)
}

func SingleResponsiblityGo() {
	user := &User{
		ID:        1,
		FirstName: "Arpan",
		LastName:  "Das",
		Email:     "dasarpan471@gmail.com",
	}
	udb := &UserDB{}
	udb.SaveUser(user)

	es := &EmailSender{}
	es.SendWelcomeEmail(user)
}
