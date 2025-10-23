package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, Go!")
// }

/* Quick Lesson

package main → required for executable programs.

func main() → program entry point.

fmt.Println → prints string with newline.

go run ./cmd/hello → compile and execute without creating binary.

*/

import "github.com/ramakrishnamamidi/go-learning/pkg/greet"

func main() {
	greet.Hello("Ramakrishna")
}

/* Quick Lesson

Packages help organize reusable code.

Only exported functions (capitalized) are accessible outside.

*/
