package main

import (
	"bufio"
	"elizaapp/doctor"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var greet string = doctor.Intro()
	fmt.Println(greet)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		// check if user the input is "quit" --> we must remove \n (enter key)
		userInput = strings.Replace(userInput, "\r\n", "", -1) // for windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // for linux and mac
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}
