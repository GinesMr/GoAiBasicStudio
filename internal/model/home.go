package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"goAiBasicStudio/internal/util"
)

var logoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00AFF0")).
	Bold(true).
	MarginBottom(1)

type home struct {
	cursor   int
	options  []string
	selected int
}

func NewHomeModel() home {
	return home{
		options:  util.ReturnOptionsMenu(),
		selected: -1,
	}
}

func (m *home) Init() tea.Cmd {
	return tea.SetWindowTitle("Home Menu")
}

func (m *home) Update(msg tea.Msg) (home, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return *m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.cursor
			if m.selected == 2 {
				return *m, func() tea.Msg {
					return showModelListMsg{}
				}
			}
		}
	}

	return *m, nil
}

func (m home) View() string {
	s := logoStyle.Render(util.AsciiLogo) + "\n"
	s += "Menu:\n\n"

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if i == m.selected {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress ↑/↓ to navigate, [Enter] to select, q to quit.\n"

	return s
}
