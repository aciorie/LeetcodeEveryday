package review

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxP(s, i, i, res)
		res = maxP(s, i, i+1, res)
	}
	return res
}
func maxP(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}
