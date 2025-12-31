package main
import "strings"

func cleanInput(text string) []string {
	strippedText := strings.TrimSpace(text)
	loweredText := strings.ToLower(strippedText)
	return strings.Fields(loweredText)
}