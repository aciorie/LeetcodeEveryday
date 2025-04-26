package leetcode75

func isSubsequence(s string, t string) bool {
	len1, len2 := len(s), len(t)
	if len1 > len2 {
		return false
	}
	i, j := 0, 0
	for i < len1 && j < len2 {
		if t[j] == s[i] {
			i++
			j++
		} else {
			j++
		}
	}
	if i != len1 {
		return false
	}
	return true
}
