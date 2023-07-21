# Dependency Inversion Principle

The dependency inversion principle states:

`High-level modules should not import anything from low-level modules. Both should depend on abstractions (e.g., interfaces).`

And,

`Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions`

Here is a code example that violates this principle:

```Golang
type User struct {
    Name         string
    Email        string
    PhoneNumber  string
    Notification Notification // Dependency on the Notification interface
}

func (u User) Notify(message string) {
    u.Notification.Send(message)
}
```

By improving

```Golang
type User struct {
    Name        string
    Email       string
    PhoneNumber string
}

func (u User) Notify(notifier Notification, message string) {
    notifier.Send(message)
}
```
