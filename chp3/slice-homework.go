package main

import "fmt"

func main() {
	var data []string
	fmt.Println(data, len(data), cap(data))

	for i := 1; i <= 10; i++ {
		value := fmt.Sprintf("%d", i)
		data = append(data, value)
	}

	fmt.Println(data)
}
