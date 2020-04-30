package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Woof", d.Name)
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Println("Woof", c.Name)
}

func main() {
	speakers := []Speaker{
		&Dog{"Bruce", true, 0},
		&Cat{"Dan", true, 0},
	}
}
