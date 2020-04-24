package main

import "fmt"

// Answer - The answer to the programming exercise
type Answer int64

const a = 2
const b float64 = 10.0
const c = 4
const answer Answer = c / a

func main() {
	fmt.Println("Constant of a kind:", a)
	fmt.Println("Constant of a type:", b)
	fmt.Println("Constant answer:", answer)

	const (
		A1 = iota
		B1 = iota
		C1 = iota
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota
		B2
		C2
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1
		B3
		C3
	)

	fmt.Println("3:", A3, B3, C3)
}
