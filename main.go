package main

import "fmt"

// reference types - типи посилань (pointers -вказівник, slices, maps, functions, channels)

// map

func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	//delete(intMap, "four")

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is in not map")
	}

	intMap["two"] = 4
	fmt.Println(intMap)

}
