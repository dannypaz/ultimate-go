package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	// person user // NOT EMBEDDING
	user  // This is embedded
	level string
}

func (u *user) notify() {
	fmt.Println("What's my name?", u.name)
}

// This overrides the notify on the notifier interface
func (a *admin) notify() {
	fmt.Println("What's my admin?", a.name)
}

func main() {
	ad := admin{
		user:  user{"Dan", "dan@gmail.com"},
		level: "super",
	}

	// embedded way
	ad.notify()

	// The non-embedded way
	ad.user.notify()

	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
