package palindrome

import "strings"

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
