package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	users := []user{
		{"ed", "ed@gmail.com"},
		{"erick", "erick@gmail.com"},
	}

	// since `u` is a copy, when we change the information..
	// it wont really matter
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	for _, u := range users {
		fmt.Println(u)
	}

	for key := range users {
		users[key].changeEmail("it@wontmatter.com")
	}

	for _, u := range users {
		fmt.Println(u)
	}
}
