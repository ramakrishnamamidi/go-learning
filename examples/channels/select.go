package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		c1 <- "from c1"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		c2 <- "from c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("Received", msg)
		case msg := <-c2:
			fmt.Println("Received", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}

/* Quick Lessons
select waits on multiple channel operations. It executes a case whose channel is ready (randomly picks if multiple ready). Use time.After in select for timeouts.
*/
