package logcheck

import (
	"slices"
	"unicode"
)

func HasCapitalLetter(str string) bool {
	if str == "" {
		return false
	}

	runes := []rune(str)
	if len(runes) == 0 {
		return false
	}
	char := runes[0]
	return unicode.IsUpper(char)
}

func HasNonEnglish(str string) bool {
	for _, v := range str {
		if unicode.IsLetter(v) && !unicode.Is(unicode.Latin, v) {
			return true
		}
	}
	return false
}

func HasSymbol(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) && !unicode.IsDigit(v) && !unicode.IsSpace(v) {
			return true
		}
	}
	return false
}

var sensitiveWords = []string{"password", "token", "secret", "apiKey"}

func HasSensitive(val string) bool {
	return slices.Contains(sensitiveWords, val)
}
