package utils

import "strings"

func TrimSuffixAndAddText(string string, suffix string, additionalText string) string {
	return strings.TrimSuffix(string, suffix) + additionalText
}
