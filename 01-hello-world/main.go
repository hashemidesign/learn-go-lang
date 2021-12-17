package main

import "fmt"

func main() {
	// var greet string = "Hello, world"
	greet := "Hello, world" // short-hand
	shout(greet)
}

func shout(sentence string) {
	fmt.Println(sentence)
}
