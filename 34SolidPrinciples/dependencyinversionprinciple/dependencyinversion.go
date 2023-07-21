package dependencyinversionprinciple

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (en EmailNotification) Send(message string) {
	fmt.Println("Sending email:", message)
}

type SMSNotification struct{}

func (sn SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// Without dependency inversion principle

// type User struct {
// 	Name         string
// 	Email        string
// 	PhoneNumber  string
// 	Notification Notification // Dependency on the Notification interface
// }

// func (u User) Notify(message string) {
// 	u.Notification.Send(message)
// }

// With the current setup, the User struct has a direct dependency on the Notification interface,
// and clients of the User struct are responsible for providing the specific implementation of Notification.

// With dependency inversion principle

type User struct {
	Name        string
	Email       string
	PhoneNumber string
}

func (u User) Notify(notifier Notification, message string) {
	notifier.Send(message)
}
func DependencyInversion() {
	user := User{Name: "John", Email: "john@example.com", PhoneNumber: "123456789"}

	emailNotifier := EmailNotification{}
	smsNotifier := SMSNotification{}

	user.Notify(emailNotifier, "Hello, this is an email notification.")
	user.Notify(smsNotifier, "Hello, this is an SMS notification.")
}
