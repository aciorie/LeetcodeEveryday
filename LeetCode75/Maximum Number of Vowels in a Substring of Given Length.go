package leetcode75

func maxVowels(s string, k int) int {
	if len(s) < k {
		return 0
	}
	vowelsInWindow := 0
	ma := make(map[byte]int, 5)
	ma['a'], ma['e'], ma['i'], ma['o'], ma['u'] = 1, 1, 1, 1, 1
	n := len(s)

	for i := 0; i < k; i++ {
		if ma[s[i]] == 1 {
			vowelsInWindow++
		}
	}
	maxVowels := vowelsInWindow

	for i := k; i < n; i++ {
		if ma[s[i]] == 1 {
			vowelsInWindow++
		}
		if ma[s[i-k]] == 1 {
			vowelsInWindow--
		}
		maxVowels = max(maxVowels, vowelsInWindow)
	}

	return maxVowels
}
