package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("My name is:", d.name)
}

// According to Bill, he does not like getter/setter methods
// because they bloat code, bloat tests, not meaningful...
// We just need to keep this in mind for later
func (d *data) setAge(age int) {
	d.age = age
}

// From Bill: Exception and not the rule is to give data behavior vs creating
// methods.
func main() {
	d := data{name: "Bill"}

	fmt.Println("Proper Calls to Methods:")

	d.displayName()
	d.setAge(45)

	fmt.Println("What the Compiler is doing:")

	// This is what Go is doing underneath
	// data.displayName(d)
	// (*data).setAge(&d, 45)

	// This is value semantics...
	f1 := d.displayName
	f1()
	// Change value of d
	d.name = "Joan"
	// Nothing happens... no change
	f1()

	// This is pointer semantics... based off of the value that
	// is used as reference point in the method setAge
	f2 := d.setAge
	d.name = "Joan"
	f2(40)
	fmt.Println(d)
}
