package pbxproj

import (
	"os"
)

func fileIsExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func isStringNeedQuote(s string) bool {
	if 0 == len(s) {
		return true
	}
	for _, c := range s {
		if c == ' ' || c == '-' || c == '?' || c == '+' || c == '@' || c == '$' || c == ',' || c == '=' ||
			c == '[' || c == ']' || c == '{' || c == '}' || c == '<' || c == '>' || c == '(' || c == ')' {
			return true
		}
	}

	return false
}
