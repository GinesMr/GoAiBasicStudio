package services

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
	models := parseWebModels(output)
	return models
}
func ListOllamaLocalModels() []string {
	cmd := exec.Command("sh", "-c", "ollama list | awk 'NR>1 {print $1}'")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error  getting local models:", err)
		return nil
	}
	output := string(out)
	localmodels := parseOllamaLocalList(output)
	return localmodels
}

func InstallNewModel(userModel string) {
	cmd := exec.Command(util.Ollama, "pull", userModel)
	msg, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(string(msg))
}
func UseSelectedModel(selectedModel string) {
	fmt.Printf("Selected model: '%s'\n", selectedModel)
	cmd := exec.Command(util.Ollama, "run", selectedModel)
	msg, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	print(string(msg))
	//USE THE SERVICE TO RUN THE MODEL
}

func runModel(selectedModel string) {

}

//TODO:Implement the function to install a new model in the local ollama but the correct way
