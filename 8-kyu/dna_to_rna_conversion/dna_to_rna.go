package kata

import "strings"

func DNAtoRNA(dna string) string {
	var sb strings.Builder
	for _, r := range dna {
		if r == 'T' {
			sb.WriteRune('U')
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
