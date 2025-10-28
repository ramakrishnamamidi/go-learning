package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func safeIncrement(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go safeIncrement(&wg)
	}
	wg.Wait()
	fmt.Println("Counter (correct):", counter)
}

/* Quick Lessons
sync.Mutex serializes access to shared variables. Lock as briefly as possible.
*/
