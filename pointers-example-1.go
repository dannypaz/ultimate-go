package main

import "fmt"

func main() {
	count := 12
	fmt.Println(count, &count)
	pointer := &count
	fmt.Println(pointer, &pointer, *pointer)
}
