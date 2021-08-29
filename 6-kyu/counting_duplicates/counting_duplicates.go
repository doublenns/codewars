package counting_duplicates

import "strings"

func duplicate_count(s1 string) int {
	chars := make(map[rune]int)
	result := 0
	for _, r := range strings.ToLower(s1) {
		_, found := chars[r]
		if !found {
			chars[r] = 1
		} else if chars[r] == 1 {
			chars[r]++
			result++
		}
	}
	return result
}
