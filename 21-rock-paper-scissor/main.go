package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')

	if playerChoice == "rock\n" {
		playerValue = ROCK
	} else if playerChoice == "paper\n" {
		playerValue = PAPER
	} else if playerChoice == "scissors\n" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
		break
	case PAPER:
		fmt.Println("Computer chose PAPER")
		break
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
		break
	default:
	}

	if playerValue == computerValue {
		fmt.Println("it's a draw")
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				fmt.Println("computer wins")
			} else {
				fmt.Println("player wins")
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				fmt.Println("computer wins")
			} else {
				fmt.Println("player wins")
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				fmt.Println("computer wins")
			} else {
				fmt.Println("player wins")
			}
			break
		default:
			fmt.Println("invalid choice")
		}
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
