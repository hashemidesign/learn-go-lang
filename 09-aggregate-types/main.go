package main

import "fmt"

// aggregate types (array, struct)
// array
var myStrings [3]string

// struct
type Car struct {
	NumberOfTires int
	Luxury        bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("First element in array is", myStrings[0])

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        false,
		Make:          "BMW",
		Model:         "X4",
		Year:          2018,
	}

	fmt.Println(myCar.Model)
}
