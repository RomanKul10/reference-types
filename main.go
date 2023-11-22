package main

import "fmt"

// reference types - типи посилань (pointers -вказівник, slices, maps, functions, channels)

// functions

func main() {
	myTotal := sumMany(2, 6, 7, 9, -5)
	fmt.Println(myTotal)
}
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total
}
