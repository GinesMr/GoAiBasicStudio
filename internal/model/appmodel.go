package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"goAiBasicStudio/internal/util"
)

type app struct {
	spinner loadingModel
}

func New() *app {
	return &app{
		spinner: newLoadingModel(),
	}
}

func (m *app) Init() tea.Cmd {
	return m.spinner.Init()
}

func (m *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case util.Quit, util.QuitC:
			return m, tea.Quit
		}
	}
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m *app) View() string {
	var show = "Getting connection to Ollama Api..." + m.spinner.spinner.View()
	return show
}
