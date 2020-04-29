package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n",
		u.name,
		u.email)
}

// When in doubt, always use reference instead of value
// structs
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	dan := user{"Danny", "dan@gmail.com"}
	dan.changeEmail("danny@gmail.com")
	dan.notify()
	fmt.Println(dan)
}
