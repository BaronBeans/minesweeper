package game

import (
	"log"
)

type Board struct {
	Width  int
	Height int
	Cells  [][]Cell
	Bombs  []Bomb
}

func NewBoard(width, height, bombcount int) Board {
	board := Board{
		Width:  width,
		Height: height,
	}
	bombs, err := GenerateBombs(width, height, bombcount)
	if err != nil {
		log.Fatal(err)
	}

	board.Bombs = bombs

	cells := GenerateCells(width, height, board)

	// PrintCells(&cells)

	board.Cells = cells

	return board
}

func (b *Board) isBomb(x, y int) bool {
	for _, bomb := range b.Bombs {
		if bomb.X == x && bomb.Y == y {
			return true
		}
	}
	return false
}
