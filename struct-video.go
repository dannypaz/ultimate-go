package main

import "fmt"

type alice struct {
	flag bool
	counter int16
	pi float32
}

type bill struct {
	flag bool
	counter int16
	pi float32
}

func main() {
	e2 := struct {
		flag bool
		counter int16
		pi float32
	}{
		flag: true,
		counter: 10,
		pi: 3.14,
	}

	var a alice
	var b bill
	b = bill(a)
	b = e2
	fmt.Println("Hello: ", a, b)
}
