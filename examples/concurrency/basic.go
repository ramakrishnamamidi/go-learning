package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, "iteration", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	// Launch a goroutine.
	go printMessage("Hello from goroutine")

	// The main goroutine continues immediately.
	printMessage("Hello from main")

	// Without waiting, program might exit before goroutine finishes.
	// Here we wait a moment to allow goroutine output to appear (not ideal; for demo only).
	time.Sleep(500 * time.Millisecond)
}
