package main

import (
	"elizaapp/doctor"
	"fmt"
)

func main() {
	var greet string = doctor.Intro()
	fmt.Println(greet)
}
