package main

import "fmt"

// add two numbers
func add(a int, b int) int {
	return a + b
}

// divide returns result and error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	sum := add(3, 4)
	fmt.Println("Sum:", sum)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

/* Quick Lesson

Functions can return multiple values.

error is the idiomatic way to handle errors in Go.

*/
