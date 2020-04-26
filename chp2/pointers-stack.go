package main

import "fmt"

const size = 1024

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, start int, stack [size]int) {
	fmt.Println(s, start)
	start++
	if start == 10 {
		return
	}
	stackCopy(s, start, stack)
}
