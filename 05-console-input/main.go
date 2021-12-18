package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	drinks := make(map[int]string)
	drinks[1] = "Coffee"
	drinks[2] = "Tea"
	drinks[3] = "Americano"
	drinks[4] = "Mocha"
	drinks[5] = "Espresso"
	drinks[6] = "Apple Juice"

	fmt.Println("Menu")
	fmt.Println("----")
	fmt.Println("1 - Coffee")
	fmt.Println("2 - Tea")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Espresso")
	fmt.Println("6 - Apple Juice")
	fmt.Println("Q - Quite the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", drinks[i]))
	}

	fmt.Println("GoodBye!...")
}
