package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 // numbers in range [2, 10]
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	guessGame(firstNumber, secondNumber, subtraction, answer)
}

func guessGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	// Display a welcome/instructions
	// take them throw the game
	// give the the answer

	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thout of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)
}
