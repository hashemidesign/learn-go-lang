package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {

	var userName string = readString("What is your name?")
	var age = readInt("How old are you?")
	// fmt.Println("Your name is "+userName+", and your age is", age, "years old")
	fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old.", userName, age))
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a value")
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number.")
		} else {
			return num
		}
	}
}
