package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type currentView int

const (
	HomeView currentView = iota
	ModelListView
	ModelLocalListView
)

type showModelListMsg struct{}
type showModelLocalListMsg struct{}

type app struct {
	view           currentView
	home           home
	modelList      newModelList
	modelLocalList newModelLocalList
}

func NewApp() *app {
	return &app{
		view:           HomeView,
		home:           NewHomeModel(),
		modelList:      NewModelList(),
		modelLocalList: NewModelLocalList(),
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

	case tea.KeyMsg:
		if msg.String() == "esc" && m.view == ModelListView {
			m.view = HomeView
			return m, nil
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
	default:
		return "Unknown view"
	}
}
