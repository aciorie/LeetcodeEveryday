package codetop4th

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	res := 0
	ma := make(map[byte]bool)
	l, r := 0, 0
	for r < len(s) {
		if ma[s[r]] {
			for ma[s[r]] {
				delete(ma, s[l])
				l++
			}
		}
		ma[s[r]] = true
		res = max(res, r-l+1)
		r++
	}
	return res
}
