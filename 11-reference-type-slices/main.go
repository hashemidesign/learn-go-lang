package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

func main() {
	// slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "horse")
	animals = append(animals, "cat")

	fmt.Println(animals)

	// Iterate through slice
	for index, x := range animals {
		fmt.Println(index, x)
	}

	// get single item
	fmt.Println(animals[2])

	// get some items
	fmt.Println(animals[0:2])

	// get slice length
	fmt.Println(len(animals))

	// sort slices
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println(animals)
	fmt.Println("Is it sorted Now?", sort.StringsAreSorted(animals))

	animals = deleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // copy last element to index i
	a[len(a)-1] = ""   // set last item to empty string
	a = a[:len(a)-1]   // truncate last item
	return a
}
