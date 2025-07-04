package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"goAiBasicStudio/internal/model"
	"os"
)

func main() {
	p := tea.NewProgram(model.New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
