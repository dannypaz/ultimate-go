package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	numbers := [4]int{10, 20, 30, 40}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	// Betty
	fmt.Printf("Bf[%s]\n", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			// Jack
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}

	// This is creating a copy of friends so even if we modify friends
	// the value will not change from the original structure....
	for i, value := range friends {
		friends[1] = "Betty"

		if i == 1 {
			// Jack (doesnt change)
			fmt.Printf("Aft[%s]\n", value)
		}
	}
}
