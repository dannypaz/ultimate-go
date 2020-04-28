package main

import "fmt"

func lengthOfArray() []string {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	return fruits
}

func capacityOfArray() []string {
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	return fruits
}

func main() {
	slice := lengthOfArray()
	inspectSlice(slice)
	slice = capacityOfArray()
	inspectSlice(slice)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p\n", i, &slice[i])

	}
}
