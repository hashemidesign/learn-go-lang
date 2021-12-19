package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func main() {
	z := addTwoNumbers(2, 5)
	fmt.Println(z)

	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println("myTotals = ", myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "bark"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}

// variatic function
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	return total
}

// type functions
func (a *Animal) Says() {
	fmt.Printf("a %s says %s\n", a.Name, a.Sound)
}
