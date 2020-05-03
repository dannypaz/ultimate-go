package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start go routines")

	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Println("waiting for go routines")
	wg.Wait()
	fmt.Println("Done with goroutines")
}

func lowercase() {
	fmt.Println("Lowercase")
}

func uppercase() {
	fmt.Println("Uppercase")
}
