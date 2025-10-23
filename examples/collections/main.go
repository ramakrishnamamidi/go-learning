package main

import "fmt"

func main() {
	// Array
	var arr = [3]int{10, 20, 30}
	fmt.Println("Array:", arr)

	// Slice
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println("Slice:", nums)

	// Map
	ages := map[string]int{"Ramakrishna": 32, "Bob": 30}
	ages["Charlie"] = 28
	fmt.Println("Map:", ages)
}

/* Quick Lesson

Arrays → fixed size.

Slices → dynamic, growable arrays.

Maps → key-value storage, unordered.

*/
