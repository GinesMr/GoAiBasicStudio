package service

import (
	"github.com/charmbracelet/glamour"
)

func MarkdownToHTML(markdown string) (string, error) {
	return glamour.Render(markdown, "dark")
}
