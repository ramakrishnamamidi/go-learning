package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Method for Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Geralt", Age: 64}
	p.Greet()
}

/* Quick Lesson

struct groups fields.

Method with value receiver (p Person) operates on copy of struct.

For modifying struct inside method, use pointer receiver (p *Person).

*/
