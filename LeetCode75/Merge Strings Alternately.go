package leetcode75

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	m, n := len(word1), len(word2)
	l, r := 0, 0
	for l < m && r < n {
		res.WriteByte(word1[l])
		res.WriteByte(word2[r])
		l++
		r++
	}
	for l < m {
		res.WriteByte(word1[l])
		l++
	}
	for r < n {
		res.WriteByte(word2[r])
		r++
	}
	return res.String()
}
