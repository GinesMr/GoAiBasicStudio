package services

import (
	"bufio"
	context2 "context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"goAiBasicStudio/internal/util"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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
	RunModel(selectedModel)
}

func RunModel(selectedModel string) string {
	bufio.NewReader(os.Stdin)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	llm, err := ollama.New(ollama.WithModel(selectedModel))
	if err != nil {
		fmt.Println("Error initializing LLM:", err)
		os.Exit(1)
	}
	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem),
	}
	cont := context2.Background()
	response := ""
	input := "This is a test input" ///Change it
	content = append(content, llms.TextParts(llms.ChatMessageTypeHuman, input))
	llm.GenerateContent(cont, content, llms.WithMaxTokens(1024),
		llms.WithStreamingFunc(func(ctx context2.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			response = response + string(chunk)
			return nil
		}))
	content = append(content, llms.TextParts(llms.ChatMessageTypeSystem, response))
	return response
}
