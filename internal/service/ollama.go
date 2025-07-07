package service

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"goAiBasicStudio/internal/util"
	"os/exec"
)

var Models []string

type OllamaFoundMsg bool

func CheckOllamaInstall() bool {
	cmd := exec.Command(util.Ollama, "--version")
	msg, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
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
func ListOllamaWebModels() []string {
	cmd := exec.Command(util.Justfile, "models")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error  models:", err)
		return nil
	}
	output := string(out)
	models := parseModels(output)
	return models
}

func InstallNewModel(userModel string) {
	cmd := exec.Command(util.Ollama, "pull", userModel)
	msg, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(string(msg))
}
