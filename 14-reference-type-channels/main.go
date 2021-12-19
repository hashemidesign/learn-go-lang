package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

func main() {
	// go routine
	go doSomething("hello, world")

	fmt.Println("This is another message")
	for {
		// do nothing
	}
}

func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is", s)
		until += 1
		if until >= 5 {
			break
		}
	}
}
