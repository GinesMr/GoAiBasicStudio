package service

import (
	"strings"
)

func parseModels(output string) []string {
	if output == "" {
		return nil
	}
	lines := strings.Split(output, "\n")
	return lines
}
