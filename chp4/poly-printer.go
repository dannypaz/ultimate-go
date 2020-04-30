package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Println("Name:", u.name)
}

func main() {
	dan := user{"Dan"}
	entities := []printer{
		dan,
		&dan,
	}

	dan.name = "Danny"

	for _, user := range entities {
		user.print()
	}

}
