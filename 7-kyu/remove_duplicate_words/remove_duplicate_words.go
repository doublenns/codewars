package removeduplicatewords

import "strings"

func RemoveDuplicateWords(str string) string {
	seen := make(map[string]bool)
	var result []string

	for _, word := range strings.Fields(str) {
		if !seen[word] {
			result = append(result, word)
			seen[word] = true
		}
	}
	return strings.Join(result, " ")
}
