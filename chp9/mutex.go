package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	const grs = 2
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < grs; count++ {
				mutex.Lock()
				{
					value := counter
					value++
					// Adding extra latency (we shouldnt do this, but used for
					// the example)
					fmt.Println(value)
					counter = value
				}
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
