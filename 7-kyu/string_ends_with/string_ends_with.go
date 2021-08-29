package string_ends_with

import "strings"

func solution1(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
