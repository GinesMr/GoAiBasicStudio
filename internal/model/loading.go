package model

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"goAiBasicStudio/internal/service"
)

type loadingModel struct {
	spinner  spinner.Model
	checking bool
	status   string
}

func NewLoadingModel() loadingModel {
	loader := spinner.New()
	loader.Spinner = spinner.Line
	loader.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true).
		PaddingLeft(1)

	return loadingModel{
		spinner:  loader,
		checking: true,
	}
}

func (model loadingModel) Init() tea.Cmd {
	return tea.Batch(
		model.spinner.Tick,
		service.CheckOllamaInstallCmd(),
	)
}

func (model *loadingModel) Update(msg tea.Msg) (loadingModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case service.OllamaFoundMsg:
		if msg {
			model.status = "✅ Ollama is installed"
		} else {
			model.status = "❌ Ollama is not installed Aborting..."
			return *model, tea.Quit
		}
		model.checking = false
	default:
		model.spinner, cmd = model.spinner.Update(msg)
	}

	return *model, cmd
}
