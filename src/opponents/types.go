package opponents

import "github.com/jbongars/connect-four-ml/src/board"

type Opponent interface {
	New(id uint8)
	MakeMove(board board.Board) (board.Board, error)
}
