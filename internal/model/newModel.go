package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"goAiBasicStudio/internal/service"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type modelsLoadedMsg []item

func ListOllamaWebModelsCmd() tea.Cmd {
	return func() tea.Msg {
		strModels := service.ListOllamaWebModels()
		items := make([]item, len(strModels))
		for i, model := range strModels {
			items[i] = item(model)
		}
		return modelsLoadedMsg(items)
	}
}

type newModelList struct {
	list     list.Model
	selected item
	loaded   bool
}

func NewModelList() newModelList {
	items := []list.Item{}
	const defaultWidth = 60
	const defaultHeight = 50
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.Title = "Available Models"
	return newModelList{
		list: l,
	}
}

func (m newModelList) Init() tea.Cmd {
	return ListOllamaWebModelsCmd()
}

func (m newModelList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case modelsLoadedMsg:
		items := make([]list.Item, len(msg))
		for i, v := range msg {
			items[i] = v
		}
		m.list.SetItems(items)
		m.loaded = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if selectedItem, ok := m.list.SelectedItem().(item); ok {
				m.selected = selectedItem
				service.InstallNewModel(string(selectedItem))
				return m, tea.Quit
			}
		case "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-2)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m newModelList) View() string {
	if !m.loaded {
		return "Getting models..\n"
	}
	return lipgloss.NewStyle().Margin(1, 2).Render(m.list.View())
}
