package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumbber  int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumbber = g.Round.RoundNumbber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)

		}
	}
}

func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) PrintIntro() {
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("----------------------")
	fmt.Println("Game is played for 3 rounds and best of three wins!")
	fmt.Println()
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumbber)
	fmt.Println("-------")
	fmt.Print("Please inter rock, paper or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))

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
		g.DisplayChan <- "it's a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.ComputerWins()
			} else {
				g.PlayerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.ComputerWins()
			} else {
				g.PlayerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.ComputerWins()
			} else {
				g.PlayerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice"
			return false
		}
	}
	return true
}

func (g *Game) ComputerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer Wins"
}

func (g *Game) PlayerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player Wins"
}

func (g *Game) PrintSummary() {
	fmt.Println("Final Score")
	fmt.Println("-----------")
	fmt.Printf("Player %d/3, Computer %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	fmt.Println()
	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins game!"
	} else {
		g.DisplayChan <- "Computer wins game!"
	}
}
