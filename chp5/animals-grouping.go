package main

import "fmt"

type animal interface {
	speak()
}

type cat struct {
	animal
	voice string
}

func (c *cat) speak() {
	fmt.Println(c.voice)
}

type dog struct {
	animal
	voice string
}

func (d *dog) speak() {
	fmt.Println(d.voice)
}

func main() {
	c := cat{voice: "Meow"}
	d := dog{voice: "Woof"}

	c.speak()
	d.speak()
}
