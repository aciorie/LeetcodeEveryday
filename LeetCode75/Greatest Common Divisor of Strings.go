package leetcode75

import "strings"

func gcdOfStrings(str1 string, str2 string) string {
	concat1, concat2 := str1+str2, str2+str1
	if concat1 != concat2 {
		return ""
	}
	
	var res strings.Builder
	m, n := len(str1), len(str2)
	gcdLen := gcd(m, n)
	for i := 0; i < gcdLen; i++ {
		res.WriteByte(str1[i])
	}
	return res.String()
}

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}
