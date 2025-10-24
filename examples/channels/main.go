package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // send value to channel
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sendNumbers(ch)

	for num := range ch {
		fmt.Println("Received:", num)
	}
}

/* Quick Lesson

Channels allow communication between goroutines.

ch <- val sends, val := <-ch receives.

range ch loops until channel is closed.

*/
