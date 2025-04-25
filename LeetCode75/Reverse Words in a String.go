package leetcode75

import (
	"strings"
)

func reverseWords(s string) string {
	trimmedS := strings.Split(strings.TrimSpace(s), " ")
	l, r := 0, len(trimmedS)-1
	for l < r {
		trimmedS[l], trimmedS[r] = trimmedS[r], trimmedS[l]
		l++
		r--
	}
	return strings.Join(trimmedS, " ")
}
