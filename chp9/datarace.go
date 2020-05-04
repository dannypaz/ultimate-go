package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < grs; count++ {
				value := counter
				value++
				// Context switch....
				fmt.Println(value)
				counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}
