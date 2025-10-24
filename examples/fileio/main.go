package main

import (
	"fmt"
	"os"
)

func main() {
	content := []byte("Hello, Go File I/O! \n")

	// Write file
	err := os.WriteFile("hello.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")

	// Read file
	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File Content:", string(data))

	// Cleanup (optional)
	_ = os.Remove("hello.txt")
}

/* Quick Lession

os.WriteFile writes bytes to file.

os.ReadFile reads all bytes from file.

Always check errors after file operations.

*/
