package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 0, "Your age")
	flag.Parse()

	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}

/* Quick Lesson

flag package parses command-line arguments.

flag.String and flag.Int define flags with default values.

Flags must be dereferenced (*name) to get actual value.


*/
