package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmc/langchaingo/llms/ollama"
)

type currentView int

const (
	HomeView currentView = iota
	ModelListView
	ModelLocalListView
	RunnerView
)

type showModelListMsg struct{}
type showModelLocalListMsg struct{}
type modelSelectedMsg string

type app struct {
	view           currentView
	home           home
	modelList      newModelList
	modelLocalList newModelLocalList
	runner         Runner
}

func NewApp() *app {
	return &app{
		view:           HomeView,
		home:           NewHomeModel(),
		modelList:      NewModelList(),
		modelLocalList: NewModelLocalList(),
		runner:         Runner{},
	}
}

func (m *app) Init() tea.Cmd {
	return m.home.Init()
}

func (m *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case showModelListMsg:
		m.view = ModelListView
		m.modelList = NewModelList()
		return m, m.modelList.Init()

	case showModelLocalListMsg:
		m.view = ModelLocalListView
		m.modelLocalList = NewModelLocalList()
		return m, m.modelLocalList.Init()

	case modelSelectedMsg:
		selected := string(msg)
		llm, err := ollama.New(ollama.WithModel(selected))
		if err != nil {
			return m, tea.Printf("Error loading model: %v", err)
		}
		m.runner = NewRunnerModel(selected, llm)
		m.view = RunnerView
		return m, m.runner.Init()

	case tea.KeyMsg:
		if msg.String() == "esc" {
			switch m.view {
			case ModelListView, ModelLocalListView, RunnerView:
				m.view = HomeView
				return m, nil
			}
		}
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}

	switch m.view {
	case HomeView:
		var cmd tea.Cmd
		m.home, cmd = m.home.Update(msg)
		return m, cmd

	case ModelListView:
		var cmd tea.Cmd
		updatedModel, cmd := m.modelList.Update(msg)
		m.modelList = updatedModel.(newModelList)
		return m, cmd

	case ModelLocalListView:
		var cmd tea.Cmd
		updatedModel, cmd := m.modelLocalList.Update(msg)
		m.modelLocalList = updatedModel.(newModelLocalList)
		return m, cmd

	case RunnerView:
		var cmd tea.Cmd
		m.runner, cmd = m.runner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *app) View() string {
	switch m.view {
	case HomeView:
		return m.home.View()
	case ModelListView:
		return m.modelList.View()
	case ModelLocalListView:
		return m.modelLocalList.View()
	case RunnerView:
		return m.runner.View()
	default:
		return "Unknown view"
	}
}
