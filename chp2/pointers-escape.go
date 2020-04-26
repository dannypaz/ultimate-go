package main

import (
	"fmt"
)

type user struct {
	name string
	email string
	age int
}

func main() {
	user1 := createUserV1()
	user2 := createUserV1()
	fmt.Println("u1", &user1, "u2", &user2)

	user3 := createUserV2()
	user4 := createUserV2()
	fmt.Println("u3", &user3, "u4", &user4)
}

func createUserV1() user {
	u := user{
		name: "Dan",
		email: "danny@amazon.com",
		age: 30,
	}
	fmt.Println("u", &u)
	return u
}

func createUserV2() *user {
	u := user{
		name: "Dan",
		email: "danny@amazon.com",
		age: 30,
	}
	fmt.Println("u", &u)
	return &u
}

