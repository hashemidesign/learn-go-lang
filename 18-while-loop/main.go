package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// simulate while loop
	i := 1000
	for i > 100 {
		// get a random numer between 1 and 1001
		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("i is", i, "so loop is going")
		}
	}
	fmt.Println("Got", i, "and broke out of loop")
}
