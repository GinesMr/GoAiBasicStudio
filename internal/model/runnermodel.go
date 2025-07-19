package model

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmc/langchaingo/llms"
	"goAiBasicStudio/internal/services"
)

type Runner struct {
	input     textarea.Model
	llm       llms.Model
	modelName string
	chat      []string
	context   []llms.MessageContent
	loading   bool
}
type RunModelMsg struct {
	Message string
}

func (r Runner) runUsingModel() tea.Cmd {
	return func() tea.Msg {
		services.RunModel(r.modelName)
		return RunModelMsg{}
	}
}
func NewRunnerModel(name string, model llms.Model) Runner {
	ta := textarea.New()
	ta.Placeholder = "Type something..."
	ta.Focus()
	return Runner{
		input:     ta,
		llm:       model,
		modelName: name,
		chat:      []string{"[System] Running model: " + name},
		context:   []llms.MessageContent{},
		loading:   false,
	}
}
func (x Runner) Init() tea.Cmd {
	return textarea.Blink
}

func (x Runner) Update(msg tea.Msg) (Runner, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return x, tea.Quit
		}
	case RunModelMsg:
		return x, tea.Quit
		//Create logic of the chat

	}
	return x, nil
}
