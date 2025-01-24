package utils

import "strings"

func ToSnakeCase(value string) string {
	toRemoveSeparators := []string{" ", ",", ".", "+", "-", "/", "\\", "@", "#", "$", "%", "^", "&", "*"}
	for _, sep := range toRemoveSeparators {
		value = strings.ReplaceAll(value, sep, "_")
	}
	return strings.ToLower(value)
}

func Tokenize(value string) []string {
	snakeValue := ToSnakeCase(value)
	return strings.Split(snakeValue, "_")
}
