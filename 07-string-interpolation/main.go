package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// create a custom type of data structure
type User struct {
	UserName        string
	Age             int
	FavourireNumber float64
	OwnADog         bool
}

func main() {
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavourireNumber = readFloat("What is your favourite number?")

	// fmt.Println("Your name is "+userName+", and your age is", age, "years old")
	fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old. Your favourite number is %.2f.",
		user.UserName,
		user.Age,
		user.FavourireNumber))
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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("Please enter a number.")
		} else {
			return num
		}
	}
}
