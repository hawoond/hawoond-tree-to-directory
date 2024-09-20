package util

import (
	"strings"
	"unicode"
)

func RemovePrefixes(line string) string {
	trimmedLine := strings.TrimLeftFunc(line, func(r rune) bool {
		return unicode.IsSpace(r) || IsSpecialChar(r)
	})
	return trimmedLine
}

func IsSpecialChar(r rune) bool {
	specialChars := []rune{'│', '├', '└', '─', '┬', '•', '+', '┝', '┞', '┡', '┢'}
	for _, c := range specialChars {
		if r == c {
			return true
		}
	}
	return false
}

func CountLeadingChars(line string) int {
	count := 0
	for _, r := range line {
		if unicode.IsSpace(r) || IsSpecialChar(r) {
			count++
		} else {
			break
		}
	}
	return count
}
