package counting_duplicates

import "strings"

func duplicate_count(s1 string) int {
	chars := make(map[rune]int)
	result := 0
	for _, r := range strings.ToLower(s1) {
		if chars[r]++; chars[r] == 2 {
			result++
		}
	}
	return result
}
