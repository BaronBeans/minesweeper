package main

import (
	"fmt"
	"minesweeper/pkg/game"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Board       game.Board
	CurrentCell game.BasicCell
}

func initialModel() model {
	return model{
		Board:       game.NewBoard(10, 10, 5),
		CurrentCell: game.BasicCell{X: 0, Y: 0},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "h":
			if m.CurrentCell.X > 0 {
				m.CurrentCell.X--
			}
		case "j":
			if m.CurrentCell.Y < len(m.Board.Cells) {
				m.CurrentCell.Y++
			}
		case "k":
			if m.CurrentCell.Y > 0 {
				m.CurrentCell.Y--
			}
		case "l":
			if m.CurrentCell.X < len(m.Board.Cells[0]) {
				m.CurrentCell.X++
			}

		case " ":
			m.Board.Cells[m.CurrentCell.Y][m.CurrentCell.X].Visible = true
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "minesweeper\n"

	for r, row := range m.Board.Cells {
		for c, col := range row {
			if m.CurrentCell.X == c && m.CurrentCell.Y == r {
				s += "_"
			} else if col.Visible == false {
				s += "*"
			} else {
				s += string(m.Board.Cells[r][c].Value)
			}
		}
		s += "\n"
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
