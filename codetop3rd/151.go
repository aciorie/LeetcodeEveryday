package codetop3rd

import "strings"

func reverseWords(s string) string {
	trimm := strings.Split(strings.TrimSpace(s), " ")
	l, r := 0, len(trimm)-1
	for l < r {
		trimm[l], trimm[r] = trimm[r], trimm[l]
		l++
		r--
	}
	return strings.Join(trimm, " ")
}
