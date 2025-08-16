package codetop

import "strings"

func reverseWords(s string) string {
	trimS := strings.Split(strings.TrimSpace(s), " ")
	l, r := 0, len(s)-1
	for l < r {
		trimS[l], trimS[r] = trimS[r], trimS[l]
		l++
		r--
	}
	return strings.Join(trimS, " ")
}
