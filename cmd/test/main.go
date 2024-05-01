package main

import (
	"minesweeper/pkg/game"
)

func main() {
	g := game.NewBoard(50, 20, 50)
	game.PrintCells(&g.Cells)
}
