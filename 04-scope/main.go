package main

import "fmt"

var one = 1 // package level varaible

func main() {
	var one = "one" // block level varaible

	fmt.Println(one)
	myFunc()

	// block level variables has priority
	// and is a bad pactice to have shadowing variables in a program
}

func myFunc() {
	// var one = "the number one"
	fmt.Println(one)
}
