package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %sWelcome to Go!\n", name)
}

/* Quick Lesson

bufio.NewReader(os.Stdin) → reads input from user.

ReadString('\n') → reads until newline.

fmt.Printf → prints formatted output.

*/
