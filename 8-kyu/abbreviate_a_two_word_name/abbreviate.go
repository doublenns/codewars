package kata

import "strings"

func AbbrevName(name string) string {
	res := make([]string, 2)
	for i, w := range strings.Fields(name) {
		res[i] = w[:1]
	}
	return strings.ToUpper(strings.Join(res, "."))
}
