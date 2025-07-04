package service

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"goAiBasicStudio/internal/util"
	"os/exec"
)

type OllamaFoundMsg bool

func CheckOllamaInstall() bool {
	cmd := exec.Command(util.Ollama, "--version")
	msg, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Printf("Ollama installation dont found")
		return false
	}
	fmt.Printf(string(msg))
	return true
}

func CheckOllamaInstallCmd() tea.Cmd {
	return func() tea.Msg {
		if CheckOllamaInstall() {
			return OllamaFoundMsg(true)
		}
		return OllamaFoundMsg(false)
	}
}
