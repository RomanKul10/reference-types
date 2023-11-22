package main

import "fmt"

// reference types - типи посилань (pointers -вказівник, slices, maps, functions, channels)

// functions

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}
func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "gaf"
	dog.NumberOfLegs = 4
	dog.says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.says()
	cat.HowManyLegs()

}
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}
	return total
}
