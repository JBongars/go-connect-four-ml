package main

import (
	"fmt"
	"os"

	"github.com/JBongars/go-connect-four-ml/src/board"
	"github.com/JBongars/go-connect-four-ml/src/opponents"
)

func pvp() {
	var board board.Board
	board.New(8)

	var player1 opponents.CliPlayerOpponent
	player1.New(1)

	var player2 opponents.CliPlayerOpponent
	player2.New(2)

	for {
		for _, elem := range []opponents.CliPlayerOpponent{player1, player2} {
			nextBoard, isWin, err := elem.MakeMove(board)

			if err != nil {
				panic(err)
			}

			if isWin {
				fmt.Printf("Player %d has won!\n", elem.GetId())
				os.Exit(0)
			}

			board = nextBoard
		}
	}
}

func main() {
	pvp()
}
