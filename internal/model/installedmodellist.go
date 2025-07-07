package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"goAiBasicStudio/internal/services"
)

type model string

func (i model) Title() string       { return string(i) }
func (i model) Description() string { return "" }
func (i model) FilterValue() string { return string(i) }

type modelsListLocalLoadedMsg []item

func ListOllamaLocalModelsCmd() tea.Cmd {
	return func() tea.Msg {
		strModels := services.ListOllamaLocalModels()
		items := make([]item, len(strModels))
		for i, model := range strModels {
			items[i] = item(model)
		}
		return modelsListLocalLoadedMsg(items)
	}
}

type newModelLocalList struct {
	list     list.Model
	selected item
	loaded   bool
}

func NewModelLocalList() newModelLocalList {
	items := []list.Item{}
	const defaultWidth = 60
	const defaultHeight = 50
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	l.Title = "Available local Models"
	return newModelLocalList{
		list: l,
	}
}

func (m newModelLocalList) Init() tea.Cmd {
	return ListOllamaLocalModelsCmd()
}

func (m newModelLocalList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				services.InstallNewModel(string(selectedItem))
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

func (m newModelLocalList) View() string {
	if !m.loaded {
		return "Getting models..\n"
	}
	return lipgloss.NewStyle().Margin(1, 2).Render(m.list.View())
}
