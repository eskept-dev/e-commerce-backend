package utils

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func RemoveDiacritics(input string) string {
	t := norm.NFD.String(input) // Decompose characters into base + diacritic
	output := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.In(r, unicode.Mn) { // Skip diacritic marks
			continue
		}
		output = append(output, r)
	}
	return string(output)
}

func ToSnakeCase(value string) string {
	value = RemoveDiacritics(value)
	toRemoveSeparators := []string{" ", ",", ".", "+", "-", "/", "\\", "@", "#", "$", "%", "^", "&", "*"}
	for _, sep := range toRemoveSeparators {
		value = strings.ReplaceAll(value, sep, "_")
	}
	return strings.ToLower(value)
}

func Tokenize(value string) []string {
	value = RemoveDiacritics(value)
	snakeValue := ToSnakeCase(value)
	return strings.Split(snakeValue, "_")
}

func NormalizeString(value string) string {
	value = RemoveDiacritics(value)
	return strings.ToLower(value)
}
