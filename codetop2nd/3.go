package codetop2nd

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	res := 0
	ma := make(map[byte]bool)
	l, r := 0, 0
	for r < n {
		if ma[s[r]] {
			for ma[s[r]] {
				delete(ma, s[l])
				l++
			}
		}
		ma[s[r]] = true
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}
	return res
}
