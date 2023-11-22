package main

import (
	"fmt"
	"sort"
)

// reference types - типи посилань (pointers -вказівник, slices, maps, functions, channels)

// slices -вказівник (specific location in memory)
func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		//i - index in the slise,
		//x - curent item in the slise that i`m looking at(поточний елемент на слайді, який я переглядаю)
		fmt.Println(i, x)
	}
	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
	fmt.Println("The slice is", len(animals), "elements long now")

}
func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // we copy last element to index i
	a[len(a)-1] = ""

	//then we erase the last element by giving it its default empty value which for a string is an empty string
	//потім ми стираємо останній елемент, надавши йому порожнє значення за замовчуванням, яке для рядка є порожнім рядком
	//and then we simply truncate the slise by deleting the last element
	// а потім ми просто обрізаємо slise, видаливши останній елемент

	a = a[:len(a)-1]
	// is len "a" will be the length of the slice minus 1 which is the last thing, so we`re returning
	// довжина "a" буде довжиною slise мінус 1, що є останнім, тому ми повертаємося

	return a
}
