package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"goAiBasicStudio/internal/service"
)

type modelsLoadedMsg []string

func ListOllamaWebModelsCmd() tea.Cmd {
	return func() tea.Msg {
		models := service.ListOllamaWebModels()
		return modelsLoadedMsg(models)
	}
}

type newModelList struct {
	cursor   int
	options  []string
	selected int
}

func NewModelList() newModelList {
	return newModelList{
		cursor:   0,
		options:  []string{},
		selected: -1,
	}
}

func (m newModelList) Init() tea.Cmd {
	return ListOllamaWebModelsCmd()
}

func (m *newModelList) Update(msg tea.Msg) (newModelList, tea.Cmd) {
	switch msg := msg.(type) {
	case modelsLoadedMsg:
		m.options = msg
		m.cursor = 0
		m.selected = -1
		return *m, nil

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
			fmt.Printf("Has seleccionado el modelo: %s\n", m.options[m.selected])
			return *m, tea.Quit
		}
	}
	return *m, nil
}

func (m newModelList) View() string {
	s := "Selecciona un modelo:\n\n"

	if len(m.options) == 0 {
		s += "Cargando modelos...\n"
	} else {
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
	}

	s += "\nUsa ↑/↓ para navegar, [Enter] para seleccionar, q para salir.\n"
	return s
}
