package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	u := user{"dan", "dan@gmail.com"}
	// Will show the original object
	fmt.Println(u)
	// Will show the String() method we created above
	fmt.Println(&u)
}
