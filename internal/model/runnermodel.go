package model

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmc/langchaingo/llms"
	"goAiBasicStudio/internal/services"
	"strings"
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

func stopModelCmd(modelName string) tea.Cmd {
	return func() tea.Msg {
		services.StopModel(modelName)
		return tea.Quit()
	}
}
func (r Runner) runUsingModel(userInput string) tea.Cmd {
	return func() tea.Msg {
		response := services.RunModel(r.modelName, userInput)
		return RunModelMsg{Message: response}
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
	var cmd tea.Cmd
	x.input, cmd = x.input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return x, stopModelCmd(x.modelName)
		case "enter":
			if !x.loading && strings.TrimSpace(x.input.Value()) != "" {
				inputText := x.input.Value()
				x.chat = append(x.chat, "[You] "+inputText)
				x.loading = true
				x.input.SetValue("")
				return x, x.runUsingModel(inputText)
			}
		}

	case RunModelMsg:
		x.chat = append(x.chat, "[AI] "+msg.Message)
		x.loading = false
	}

	return x, cmd
}
func (r Runner) View() string {
	output := strings.Join(r.chat, "\n")
	return fmt.Sprintf(
		"%s\n\n%s\n\n[Enter: send | q: exit]",
		output,
		r.input.View(),
	)
}
