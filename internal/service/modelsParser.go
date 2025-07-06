package service

import "strings"

func parseModels(output []byte) []string {

	//TODO: Check if the output is empty or contains only whitespace
	parts := string(output)
	if parts == "" {
		return nil
	}
	lines := strings.Split(parts, "\n")
	return lines
}
