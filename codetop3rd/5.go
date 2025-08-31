package codetop3rd

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxp(s, i, i, res)
		res = maxp(s, i, i+1, res)
	}
	return res
}

func maxp(s string, l, r int, res string) string {
	sub := ""
	for l >= 0 && r < len(s) && s[l] == s[r] {
		sub = s[l : r+1]
		l--
		r++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
