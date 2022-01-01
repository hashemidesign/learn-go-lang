package main

import "fmt"

func main() {
	a := 12
	b := 6

	// notice the &&, left hand must be true to check second hand
	if b != 0 && divideTwoNumbers(a, b) == 2 {
		fmt.Println("found a two")
	}
}

func divideTwoNumbers(a, b int) int {
	return a / b
}
