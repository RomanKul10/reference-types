package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// reference types - типи посилань (pointers -вказівник, slices, maps, functions, channels)

// channels
var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
