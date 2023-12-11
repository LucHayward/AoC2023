package utils

import "strings"

// splitLines normalizes line endings to UNIX style and splits the input into lines.
func SplitLines(input string) []string {
	normalizedInput := strings.Replace(input, "\r\n", "\n", -1)
	return strings.Split(normalizedInput, "\n")
}
