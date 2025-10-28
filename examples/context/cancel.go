package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker: stopped:", ctx.Err())
			return
		default:
			fmt.Println("worker: working")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(time.Second)
	fmt.Println("main: cancelling")
	cancel()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("main: done")
}

/* Quick Lessons
context carries cancellation, deadlines, and request-scoped values. Pass contexts to goroutines to allow cooperative cancellation.
*/
