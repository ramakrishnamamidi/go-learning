package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "ping"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)
}

/* Quick Lessons
Channels are typed conduits that synchronize send/receive operations. An unbuffered channel blocks the sender until a receiver is ready.
*/
