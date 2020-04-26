package main

import (
	"fmt"
)

func main() {
	count := 1
	fmt.Println("Value:", count, "Memory:", &count)
	incr(&count)
	fmt.Println("Value:", count, "Memory:", &count)
}

func incr(count *int) {
	*count++
	fmt.Println(
		count, 	// The actual pointer.. memory location
		&count, // The pointer to the pointer
		*count, // the value
	)
	fmt.Println("Value:", count, "Memory:", &count)
}
