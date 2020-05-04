package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var rwMutex sync.RWMutex

var data []string

// Number of reads occuring at any given time
var readCount int64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	for i := 0; i < 8; i++ {
		go func() {
			for {
				reader(i)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}

func writer(idx int) {
	rwMutex.Lock()
	{
		rc := atomic.LoadInt64(&readCount)
		fmt.Printf("Yolo", rc)
		data = append(data, fmt.Sprintf("String %d", idx))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	rwMutex.RLock()
	{
		rc := atomic.AddInt64(&readCount, 1)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Println(id, len(data), rc)
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}
