package utils

import (
	"strings"
)

func ToCapitalize(str string) string {
	return strings.ToUpper(str[0:1]) + strings.ToLower(str[1:])
}
