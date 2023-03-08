package opponents

import (
	"fmt"

	"github.com/jbongars/connect-four-ml/src/board"
)

type CliPlayerOpponent struct {
	id uint8
}

func (o *CliPlayerOpponent) New(id uint8) {
	o.id = id
}

func (o *CliPlayerOpponent) GetId() uint8 {
	return o.id
}

func (o *CliPlayerOpponent) MakeMove(board board.Board) (board.Board, bool, error) {
	var position int

	fmt.Print("\n\n")

	board.Print()
	fmt.Println("---------------------------------------")
	fmt.Printf("Move (0 - %d): ", board.GetSize())
	fmt.Scanln(&position)

	if !board.IsValidMove(position) {
		fmt.Printf("Invalid move!")
		return o.MakeMove(board)
	}

	nextBoard, isWin, err := board.MakeMove(o.id, position)

	if err != nil {
		return nextBoard, isWin, err
	}

	return nextBoard, isWin, nil
}
