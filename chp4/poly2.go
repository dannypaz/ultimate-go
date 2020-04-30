package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println(u.name, u.email)
}

func main() {
	dan := user{"Dan", "dan@gmail.com"}
	sendNotification(&dan)
}

func sendNotification(n notifier) {
	n.notify()
}
