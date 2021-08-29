package remove_string_spaces

import "strings"

func NoSpace1(word string) string {
	var result strings.Builder
	for _, v := range word {
		if v != ' ' {
			result.WriteRune(v)
		}
	}
	return result.String()
}
