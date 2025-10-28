package main

import (
	"fmt"
	"time"
)

var counter = 0

func increment() {
	// This increments shared counter without synchronization -> race condition
	counter++
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	// Give goroutines time to run (for demo only)
	time.Sleep(1 * time.Second)
	fmt.Println("Final Counter (likely incorrect):", counter)
}

/* Quick Lesson
Race condition: multiple goroutines access/modify shared variable without synchronization — leads to nondeterministic results.

Use the race detector tool -race to find races at runtime.

*/
