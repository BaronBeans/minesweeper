package main

import (
	"fmt"
	"minesweeper/pkg/game"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Board       game.Board
	CurrentCell game.BasicCell
	GameOver    bool
}

func initialModel() model {
	return model{
		Board:       game.NewBoard(50, 20, 100),
		CurrentCell: game.BasicCell{X: 0, Y: 0},
		GameOver:    false,
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
				return m, nil
			}
		case "j":
			if m.CurrentCell.Y < len(m.Board.Cells)-1 {
				m.CurrentCell.Y++
				return m, nil
			}
		case "k":
			if m.CurrentCell.Y > 0 {
				m.CurrentCell.Y--
				return m, nil
			}
		case "l":
			if m.CurrentCell.X < len(m.Board.Cells[0])-1 {
				m.CurrentCell.X++
				return m, nil
			}

		case " ":
			m.Board.HitCell(m.CurrentCell.X, m.CurrentCell.Y)
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.GameOver {
		s := "Game over, won!!!"
		return s
	}

	s := "minesweeper\n"

	for r, row := range m.Board.Cells {
		for c, col := range row {
			if m.CurrentCell.X == c && m.CurrentCell.Y == r {
				s += "_"
			} else if col.Visible == false {
				s += "*"
			} else if col.Value == -1 {
				s += "!"
			} else {
				s += fmt.Sprintf("%d", m.Board.Cells[r][c].Value)
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
