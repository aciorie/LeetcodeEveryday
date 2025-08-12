package slidingwindow

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	res, ma := 0, make(map[byte]bool)
	l, r := 0, 0

	for r < len(s) {
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
