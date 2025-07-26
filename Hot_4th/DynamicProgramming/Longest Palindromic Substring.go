package dynamicprogramming

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1, len2 := expand(s, i, i), expand(s, i, i+1)
		curMaxLen := max(len1, len2)

		if curMaxLen > end-start+1 {
			// 例如，如果长度是 5，中心是 i，那么它的范围是 [i - 2, i + 2]。
            // 此时 (i - (currentMaxLen - 1) / 2)
            // 例如，长度是 4，中心是 i, i+1，那么它的范围是 [i - 1, i + 2]。
            // 此时 (i - (currentMaxLen / 2) + 1)
            // 综合一下就是：start = i - (currentMaxLen - 1) / 2
			start = i - (curMaxLen-1)/2
			end = i + curMaxLen/2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) int {
	// 只要左右指针不越界，且指向的字符相同，就继续扩展
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 循环结束时，left 和 right 已经越界或字符不匹配
	// 回文串的实际长度是 (right - 1) - (left + 1) + 1 = right - left - 1
	return right - left - 1
}
