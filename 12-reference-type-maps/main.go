package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

func main() {
	// maps
	// to create a map we have to **make it
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// maps are not sorted!! and not sortable!
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// remove item from map
	delete(intMap, "four")
	fmt.Println(intMap)

	// element exist in a map
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in the map")
	} else {
		fmt.Println(el, "is innot in the map")
	}

	// change value
	intMap["two"] = 22
	fmt.Println(intMap)
}
