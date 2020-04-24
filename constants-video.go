package main

import "fmt"

// Duration of our consts
type Duration int64

// Duration of the stuff we are doing
const (
	Nanosecond Duration = 1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second = 1000 * Millisecond
	Minute = 60 * Second
	Hour = 60 * Minute
)

func main() {
	d := add(Nanosecond)
	fmt.Println(d)
}

func add(d Duration) Duration {
	return d + 1
}
