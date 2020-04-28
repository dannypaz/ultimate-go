package main

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

func map1() {
	users := make(map[string]user)
	users["Dan"] = user{"Dan", "Paz"}
	users["Jessie"] = user{"Jessie", "Sheng"}

	for key, val := range users {
		fmt.Println(key, val)
	}

	fmt.Println()

	for key := range users {
		fmt.Println(key)
	}
}

func map2() {
	users := map[string]user{
		"Dan":    {"Dan", "Paz"},
		"Jessie": {"Jessie", "Sheng"},
		"Roy":    {"Roy", "Lee"},
	}

	for key, value := range users {
		fmt.Println(key, value)
	}

	delete(users, "Roy")

	fmt.Println()

	u, found := users["Roy"]

	fmt.Println("Roy", found, u)
}

func main() {
	users := map[string]user{
		"Roy":    {"Roy", "Lee"},
		"Dan":    {"Dan", "Paz"},
		"Andy":   {"Andy", "Paz"},
		"Jessie": {"Jessie", "Sheng"},
	}

	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}
