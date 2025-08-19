package codetop2nd

func longestPalindrome(s string) string {
	n := len(s)
	res := ""
	for i := 0; i < n; i++ {
		res = maxp(s, i, i, res)
		res = maxp(s, i, i+1, res)
	}
	return res
}

func maxp(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j <= len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
