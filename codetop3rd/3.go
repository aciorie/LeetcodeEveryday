package codetop3rd

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	res, seen := 0, make(map[byte]bool)
	l, r := 0, 0
	for r < n {
		if seen[s[r]] {
			for seen[s[r]] {
				delete(seen, s[l])
				l++
			}
		}

		seen[s[r]] = true
		res = max(res, r-l+1)
		r++
	}
	return res
}
