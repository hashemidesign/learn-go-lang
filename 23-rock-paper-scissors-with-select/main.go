package main

import (
	"myapp/game"
)

func main() {

	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumbber:  0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	go game.Rounds()

	game.ClearScreen()
	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan

		if game.Round.RoundNumbber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}

	game.PrintSummary()
}
