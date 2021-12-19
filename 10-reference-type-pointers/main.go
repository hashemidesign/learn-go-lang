package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

func main() {
	// pointers

	x := 10
	myFirstPointer := &x
	fmt.Println("x is ", x)
	fmt.Println("myFirstPointer is ", myFirstPointer)
	// * is pointer indicator
	*myFirstPointer = 15
	fmt.Println("x is ", x)

	changeValueOfPointer(myFirstPointer)
	fmt.Println("x is ", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
