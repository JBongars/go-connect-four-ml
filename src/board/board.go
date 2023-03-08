package board

import (
	"errors"
	"fmt"
)

type Board struct {
	state []uint8
	size  int
	win   int
}

func (b *Board) New(size int) {
	b.size = size
	b.state = make([]uint8, int(size)*int(size))
	b.win = 4
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) getRow(row uint8) []uint8 {
	rowStart := int(row) * int(b.size)
	result := make([]uint8, int(b.size))

	for i := 0; i < int(b.size); i++ {
		result[i] = b.state[i+rowStart]
	}
	return result
}

func (b *Board) getColumn(col uint8) []uint8 {
	result := make([]uint8, b.size)

	for i := 0; i < int(b.size); i++ {
		columnIndex := (i * int(b.size)) + int(col)
		result[i] = b.state[columnIndex]
	}
	return result
}

func (b *Board) IsValidMove(position int) bool {
	if position > b.GetSize()-1 {
		return false
	}

	for i, elem := range b.state {
		if i%int(b.size) == int(position) && elem == 0 {
			return true
		}
	}
	return false
}

func (b *Board) CheckWin() uint8 {
	for p := uint8(1); p <= 2; p++ {
		for r := 0; r < b.size; r++ {
			for c := 0; c < b.size; c++ {
				idx := r*b.size + c
				if b.state[idx] != p {
					continue
				}
				for _, dir := range []struct{ dr, dc int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}} {
					// check if there is enough space in this direction
					if dir.dc < 0 && c < b.win-1 || dir.dc == 0 && c < b.win-1 || dir.dc > 0 && c > b.size-b.win {
						continue
					}
					if dir.dr < 0 && r < b.win-1 || dir.dr > 0 && r > b.size-b.win {
						continue
					}
					// check for b.win in this direction
					match := true
					for i := 1; i < b.win; i++ {
						r2 := r + (i)*(dir.dr)
						c2 := c + (i)*(dir.dc)
						if b.state[r2*b.size+c2] != p {
							match = false
							break
						}
					}
					if match {
						return p
					}
				}
			}
		}
	}
	return 0
}

func (b *Board) CheckWinAtIndex(idx int) bool {
	rowIdx, colIdx := idx/b.size, idx%b.size
	player := b.state[idx]

	for _, dir := range []struct{ dRow, dCol int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}} {
		count := 0
		for i := 0; i < int(b.win); i++ {
			nextRow := rowIdx + i*dir.dRow
			nextCol := colIdx + i*dir.dCol
			if nextRow >= b.size || nextCol >= b.size || nextRow < 0 || nextCol < 0 {
				break
			}
			if b.state[nextRow*b.size+nextCol] == player {
				count++
			} else {
				break
			}
		}
		if count == int(b.win) {
			return true
		}
	}
	return false
}

func (b *Board) MakeMove(player uint8, position int) (Board, bool, error) {
	for i, elem := range b.state {
		if i%int(b.size) == int(position) && elem == 0 {
			nextB := b.Clone()
			nextB.state[i] = player
			nextB.Print()
			win := nextB.CheckWinAtIndex(i)

			return nextB, win, nil
		}
	}
	return Board{}, false, errors.New("invalid move")
}

func (b *Board) Clone() Board {
	clone := Board{size: b.size, win: b.win, state: make([]uint8, len(b.state))}

	for i, elem := range b.state {
		clone.state[i] = elem
	}
	return clone
}

func (b *Board) Print() {
	// Define a list of symbols to use for each player
	symbols := []string{"⚪️", "🔴", "🔵", "🟢", "🟡", "🟣", "🟤", "⚫️"}

	// Print each row
	for r := int(b.size) - 1; r >= 0; r-- {
		fmt.Print("| ")
		for c := 0; c < b.size; c++ {
			idx := r*int(b.size) + int(c)
			if b.state[idx] >= uint8(len(symbols)) {
				fmt.Print("  ")
			} else {
				fmt.Print(symbols[b.state[idx]] + " ")
			}
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
}

func (b *Board) DebugPrint() {
	for r := int(b.size) - 1; r >= 0; r-- {
		fmt.Print("| ")
		for c := 0; c < b.size; c++ {
			idx := r*int(b.size) + int(c)
			fmt.Printf("%d |", b.state[idx])
		}
		fmt.Print("\n")
	}
}
