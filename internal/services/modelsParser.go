package services

import (
	"strings"
)

var ModelLocalList = []string{}

func parseWebModels(output string) []string {
	if output == "" {
		return nil
	}
	lines := strings.Split(output, "\n")
	return lines
}

func parseOllamaLocalList(output string) []string {
	if output == "" {
		return nil
	}
	names := strings.Split(output, "\n")
	return names
}
