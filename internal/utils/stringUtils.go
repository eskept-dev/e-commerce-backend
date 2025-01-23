package utils

import "strings"

func ToSnakeCase(value string) string {
	return strings.ToLower(strings.ReplaceAll(value, " ", "_"))
}
