package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// reference types (pointers, slices, maps, functions, channels)

// you can only pass 1 type to a channel
// rune is a single character
var keyPressChan chan rune

func main() {
	// we have to **make the channel
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("press any key, or q to quit")
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
		fmt.Println("you pressed", string(key))
	}
}
