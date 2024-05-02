package main

import (
	"fmt"
	"minesweeper/pkg/game"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	Board       game.Board
	CurrentCell game.BasicCell
	GameOver    bool
}

func initialModel() model {
	return model{
		Board:       game.NewBoard(10, 10, 10),
		CurrentCell: game.BasicCell{X: 0, Y: 0},
		GameOver:    false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Board.GameState == "won" || m.Board.GameState == "lost" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
			return m, nil
		}
	}

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
	s := "minesweeper\n"

	baseStyle := lipgloss.NewStyle().Padding(0, 1)

	for r, row := range m.Board.Cells {
		for c, col := range row {
			visible := m.Board.Cells[r][c].Visible
			if m.CurrentCell.X == c && m.CurrentCell.Y == r {
				selectedStyle := lipgloss.NewStyle().Background(lipgloss.Color("#ffffff")).Foreground(lipgloss.Color("#9c00ff")).Padding(0, 1)
				if visible {
					if col.Value == -1 {
						s += selectedStyle.Render("!")
					} else if col.Value == 0 {
						s += selectedStyle.Render(".")
					} else {
						s += selectedStyle.Render(fmt.Sprintf("%d", m.Board.Cells[r][c].Value))
					}
				} else {
					s += selectedStyle.Render("*")
				}
			} else if visible {
				if col.Value == -1 {
					s += baseStyle.Render("!")
				} else if col.Value == 0 {
					s += baseStyle.Render(".")
				} else {
					s += baseStyle.Render(fmt.Sprintf("%d", m.Board.Cells[r][c].Value))
				}

			} else {
				s += baseStyle.Render("*")
			}
		}
		s += "\n"
	}

	if m.Board.GameState == "won" {
		s += "You won!"
	} else if m.Board.GameState == "lost" {
		s += "You lost!"
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
