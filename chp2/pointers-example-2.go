package main

import "fmt"

type jessie struct {
	name string
	isCool bool
}

func main() {
	person := jessie{name: "Dan", isCool: true}
	fmt.Println(person)
	changeName(&person)
	fmt.Println(person)
}

func changeName(person *jessie) {
	person.name = "Jessie"
}
