package main

import "myapp/packageone"

var myVar = "package var in main"

func main() {
	var blockVar = "block var in main"
	packageone.PrintMe(blockVar, myVar)
}
