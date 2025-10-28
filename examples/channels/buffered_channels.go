package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // would block until a receiver consumes

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*  Quick Lessons
Buffered channels decouple producer/consumer: senders can send up to buffer capacity without blocking.
*/
