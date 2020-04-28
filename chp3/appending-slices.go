package main

import "fmt"

func emptyStruct() {
	var data []string
	// empty struct, this creates a pointer and this pointer
	// is pointing to a global struct{} empty value so we can
	// continue to create empty structs but there would never
	// be an allocation of memory
	data = []string{}
	// Empty struct example which is a special type in go
	var es struct{}

	fmt.Println(data, es)
}

func main() {
	// nil slice
	var data []string

	lastcap := cap(data)

	fmt.Println(lastcap, data, 1e5)

	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)
		break
	}
}
