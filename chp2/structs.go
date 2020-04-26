package main

import (
	"fmt"
)

type Person struct {
	name  string
	email string
	age   int
}

func main() {
	peer := Person{
		name:  "Danny",
		email: "danny@gmail.com",
		age:   10,
	}

	fmt.Println(peer)

	person2 := struct{
		name  string
		email string
		age   int
	}{
		name:  "Art",
		email: "art@gmail.com",
		age:   20,
	}

	fmt.Println(person2)

}
