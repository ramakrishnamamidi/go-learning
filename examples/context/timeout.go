package main

import (
	"context"
	"fmt"
	"time"
)

func slowOp(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("slowOp: done")
	case <-ctx.Done():
		fmt.Println("slowOp: cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	slowOp(ctx)
}

/* Quick Lessons
context.WithTimeout cancels automatically after the duration. Useful to bound long operations.
*/
