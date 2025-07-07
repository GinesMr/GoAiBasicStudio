package services

import (
	"bufio"
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

func parselocalModels(output string) []string {
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

	var names []string
	scanner := bufio.NewScanner(strings.NewReader(output))
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		if lineNumber == 0 {
			lineNumber++
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 1 {
			names = append(names, fields[0])
		}
	}
	return names
}
