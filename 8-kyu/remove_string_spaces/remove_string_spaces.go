package remove_string_spaces

import (
	"strings"
	"unicode"
)

func NoSpace1(word string) string {
	var result strings.Builder
	for _, v := range word {
		if v != ' ' {
			result.WriteRune(v)
		}
	}
	return result.String()
}

func NoSpace2(word string) string {
	var result strings.Builder
	for _, v := range word {
		if !unicode.IsSpace(v) {
			result.WriteRune(v)
		}
	}
	return result.String()
}

func NoSpace3(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func NoSpace4(word string) string {
	ws := strings.Fields(word)
	return strings.Join(ws, "")
}
