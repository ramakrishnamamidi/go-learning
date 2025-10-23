package main

import "fmt"

func main() {
	// Explicit type
	var age int = 32
	var name string = "Ramakrishna"

	// Implicit type (inference)
	city := "Hyderabad"

	// Constant
	const country = "India"

	fmt.Printf("%s is %d years old and lives in %s, %s.\n", name, age, city, country)
}

/* Quick Lesson

var → declares variable.

:= → shorthand with type inference.

const → immutable value.

fmt.Printf → formatted output.

*/
