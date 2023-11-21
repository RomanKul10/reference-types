package main

import "fmt"

// reference types - типи посилань (pointers -вказівник, slices, functions, channels)

// pointers -вказівник (specific location in memory)
func main() {
	x := 10

	myFirstPointer := &x // address the pointer to the value of 10

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15 // address in memory where  first pointer is pointing
	fmt.Println("x is now", x)

	changeValueofPointer(&x)
	fmt.Println("After function call, x is now", x)

}
func changeValueofPointer(num *int) {
	*num = 25
}
