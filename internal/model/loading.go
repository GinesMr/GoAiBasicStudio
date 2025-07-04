package model

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type loadingModel struct {
	spinner spinner.Model
}

func newLoadingModel() loadingModel {
	loader := spinner.New()
	loader.Spinner = spinner.Line
	loader.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true).
		PaddingLeft(1)

	return loadingModel{
		spinner: loader,
	}
}

func (model loadingModel) Init() tea.Cmd {
	return model.spinner.Tick
}

func (model *loadingModel) Update(msg tea.Msg) (loadingModel, tea.Cmd) {
	var cmd tea.Cmd
	model.spinner, cmd = model.spinner.Update(msg)
	return *model, cmd
}
