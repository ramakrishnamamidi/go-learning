package main

import "fmt"

func increment(a *int) {
	*a = *a + 1
}

func main() {
	x := 10
	fmt.Println("Before:", x)
	increment(&x)
	fmt.Println("After:", x)

	// Pointer with struct
	type Person struct {
		Name string
		Age  int
	}

	p := &Person{Name: "Bob", Age: 25}
	fmt.Println("Person:", *p)
	p.Age += 5
	fmt.Println("Updated Person:", *p)
}

/* Quick Lessons

*int → pointer to int.

&x → address of variable x.

*p → value at pointer p.

Pointer receivers allow methods to modify original struct


| Concept              | Code             | Meaning                               |
| -------------------- | ---------------- | ------------------------------------- |
| Pointer to int       | `*int`           | stores address of int                 |
| Address of variable  | `&x`             | gets memory address                   |
| Value at pointer     | `*p`             | dereferences pointer                  |
| Struct pointer print | `fmt.Println(p)` | Go shows `&{...}`                     |
| Struct field access  | `p.Age`          | Go automatically dereferences pointer |


*/
