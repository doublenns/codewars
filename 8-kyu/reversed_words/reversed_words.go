package reversed_words

import (
	"strings"
)

func ReverseWords(str string) string {
	ss := strings.Fields(str)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return strings.Join(ss, " ")
}
