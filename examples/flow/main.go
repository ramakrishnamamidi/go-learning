package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend soon!")
	default:
		fmt.Println("Midweek")
	}
}

/* Quick Lesson

Go only has for loops (no while).

switch does not fall through by default.

if conditions must be boolean expressions.

*/
