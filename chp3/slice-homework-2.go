package main

import "fmt"

func inspectSlice(slice []string) {
	for i := range slice {
		fmt.Printf("Index[%d] value[%v] pointer[%v]\n", i+1, slice[i], &slice[i])
	}
}

func main() {
	data := make([]string, 5)
	data[0] = "Danny"
	data[1] = "Kevin"
	data[2] = "Josh"
	data[3] = "Mo"
	data[4] = "Jeff"

	inspectSlice(data)
	slice1 := data[1:3]
	inspectSlice(slice1)
}
