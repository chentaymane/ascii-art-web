package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Run(input string, banner string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("400: empty input")
	}

	if len(input) > 2000 {
		return "", fmt.Errorf("400: invalid size (max 2000 chars)")
	}

	// Allowed banners
	allowed := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !allowed[banner] {
		return "", fmt.Errorf("404: banner not found")
	}

	fontPath := "banners/" + banner + ".txt"

	// Normalize line breaks
	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	content, err := os.ReadFile(fontPath)
	if err != nil {
		return "", fmt.Errorf("404: banner file missing")
	}

	fontTxt := strings.ReplaceAll(string(content), "\r\n", "\n")
	fontLines := strings.Split(fontTxt, "\n")

	if isOnlyNewline(input) {
		return input, nil
	}
	final := ""

	for _, line := range lines {

		if line == "" {
			final += "\n"
			continue
		}
		for row := 1; row < 9; row++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", fmt.Errorf("400: invalid character")
				}

				index := int(char-32)*9 + row

				if index < 0 || index >= len(fontLines) {
					return "", fmt.Errorf("500: corrupted banner file")
				}

				final += fontLines[index]
			}
			final += "\n"
		}
	}

	return final, nil
}

func isOnlyNewline(s string) bool {
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	return true
}
