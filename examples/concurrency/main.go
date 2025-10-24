package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from goroutine")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello() // start goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from main")
		time.Sleep(150 * time.Millisecond)
	}
}

/* Quick Lesson

go keyword launches a goroutine (lightweight thread).

Goroutines run concurrently with main program.

Use time.Sleep to prevent main from exiting early.

*/
