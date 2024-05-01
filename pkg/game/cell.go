package game

type BasicCell struct {
	X, Y int
}

type Cell struct {
	X       int
	Y       int
	Value   int // -1 for mine, 0 for empty, 1-8 for number of surrounding mines
	Visible bool
}

func GenerateCells(width, height int, board Board) [][]Cell {
	cells := make([][]Cell, 0, height)
	for y := 0; y < height; y++ {
		row := make([]Cell, 0, width)
		for x := 0; x < width; x++ {
			if board.isBomb(x, y) {
				row = append(row, Cell{
					X:       x,
					Y:       y,
					Value:   -1,
					Visible: false,
				})
			} else {
				bombCount := countSurround(x, y, board)
				row = append(row, Cell{
					X:       x,
					Y:       y,
					Value:   bombCount,
					Visible: false,
				})
			}

		}
		cells = append(cells, row)
	}
	return cells
}

func countSurround(x, y int, board Board) int {
	dirs := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	bombCount := 0
	for _, dir := range dirs {
		// newPoint := BasicCell{X: x + dir[1], Y: y + dir[0]}
		b := board.isBomb(x+dir[1], y+dir[0])
		// fmt.Println(newPoint, b)
		if b {
			bombCount++
		}
	}

	// fmt.Println("bombCount", bombCount)
	return bombCount
}

func PrintCells(cells *[][]Cell) {
	for y := 0; y < len(*cells); y++ {
		for x := 0; x < len((*cells)[y]); x++ {
			if (*cells)[y][x].Value == -1 {
				print("X")
			} else {
				print((*cells)[y][x].Value)
			}
		}
		println()
	}
}
