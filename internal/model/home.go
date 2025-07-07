package model

import (
	"fmt"
	"goAiBasicStudio/internal/service"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"goAiBasicStudio/internal/util"
)

type markdownRenderedMsg string

func loadMarkdownCmd(md string) tea.Cmd {
	return func() tea.Msg {
		rendered, err := service.MarkdownToHTML(md)
		if err != nil {
			rendered = "Error rendering markdown: " + err.Error()
		}
		return markdownRenderedMsg(rendered)
	}
}

var logoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00AFF0")).
	Bold(true).
	MarginBottom(1)

type home struct {
	cursor     int
	options    []string
	selected   int
	outputText string
}

func NewHomeModel() home {
	return home{
		options:    util.ReturnOptionsMenu(),
		selected:   -1,
		outputText: "",
	}
}

func (m *home) Init() tea.Cmd {
	return tea.SetWindowTitle("Home Menu")
}

func (m *home) Update(msg tea.Msg) (home, tea.Cmd) {
	switch msg := msg.(type) {

	case markdownRenderedMsg:
		m.outputText = string(msg)
		return *m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return *m, tea.Quit

		case "esc":
			if m.outputText != "" {
				m.outputText = ""
				return *m, nil
			}

		case "up", "k":
			if m.outputText == "" && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.outputText == "" && m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.outputText == "" {
				m.selected = m.cursor

				if m.selected == 2 {
					return *m, func() tea.Msg {
						return showModelListMsg{}
					}
				}
				if m.selected == 3 {
					md := util.Md
					return *m, loadMarkdownCmd(md)
				}
			}
		}
	}

	return *m, nil
}

func (m home) View() string {
	if m.outputText != "" {
		return m.outputText + "\n\n[ESC] para volver"
	}

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
