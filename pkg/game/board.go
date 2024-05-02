package game

import (
	"log"
)

type Board struct {
	Width     int
	Height    int
	Cells     [][]Cell
	Bombs     []Bomb
	GameState string
}

func NewBoard(width, height, bombcount int) Board {
	board := Board{
		Width:     width,
		Height:    height,
		GameState: "playing",
	}
	bombs, err := GenerateBombs(width, height, bombcount)
	if err != nil {
		log.Fatal(err)
	}

	board.Bombs = bombs

	cells := GenerateCells(width, height, board)

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

func (b *Board) HitCell(x, y int) {
	b.Cells[y][x].Visible = true
	if b.Cells[y][x].Value == 0 {
		b.checkZeroNeighbors(y, x)
	}

	won := b.checkWinState()
	if won {
		b.GameState = "won"
	}

	lost := b.checkLoseState()
	if lost {
		b.GameState = "lost"
	}
}

func (b *Board) checkZeroNeighbors(row int, col int) {
	// Define the possible directions to move in the matrix
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	// Check all 4 directions around the current cell
	for _, direction := range directions {
		newRow, newCol := row+direction[1], col+direction[0]

		// Check if the new position is within the matrix boundaries
		if newRow >= 0 && newRow < len(b.Cells) && newCol >= 0 && newCol < len(b.Cells[0]) {
			// If the neighbor is 0 and it hasn't been visited yet, return true
			if b.Cells[newRow][newCol].Value == 0 && !b.Cells[newRow][newCol].Visible {
				// Mark the current cell as visited
				b.Cells[row][col].Visible = true
				b.checkZeroNeighbors(newRow, newCol)
			}
		}
	}
}

func (b *Board) checkWinState() bool {
	var invisibleCells []BasicCell
	for _, row := range b.Cells {
		for _, col := range row {
			if !col.Visible {
				invisibleCells = append(invisibleCells, BasicCell{X: col.X, Y: col.Y})
			}
		}
	}

	if len(invisibleCells) == len(b.Bombs) {
		return true
	}

	return false
}

func (b *Board) checkLoseState() bool {
	for _, row := range b.Cells {
		for _, col := range row {
			if b.isBomb(col.X, col.Y) && col.Visible {
				return true
			}
		}
	}
	return false
}
