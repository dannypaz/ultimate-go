package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	const grs = 2
	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < grs; count++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	fmt.Println("What happened")
	wg.Wait()
	fmt.Println(counter)
}
