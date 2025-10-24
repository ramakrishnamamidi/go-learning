package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	// Block until all workers have called Done().
	wg.Wait()
	fmt.Println("All workers completed")
}

/* Quick Lesson
sync.WaitGroup lets the main goroutine wait for a collection of goroutines to finish.

Use wg.Add(n) before launching goroutines, and each goroutine calls defer wg.Done()
*/
