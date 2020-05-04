package main

import "sync"

var scores = make(map[string]int)

// This will fail to show map datarace
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()
	wg.Wait()
}
