package codetop2nd

import "strings"

func reverseWords(s string) string {
	trimmed := strings.Split(strings.TrimSpace(s), " ")
	l, r := 0, len(trimmed)-1
	for l < r {
		trimmed[l], trimmed[r] = trimmed[r], trimmed[l]
		l++
		r--
	}
	return strings.Join(trimmed, " ")
}
