package hot2nd

func partition(s string) [][]string {
	if len(s) == 1 {
		return [][]string{{string(s[0])}}
	}
	var res [][]string
	var cur []string
	var dfs func(s string, cur []string)
	dfs = func(s string, cur []string) {
		if len(s) == 0 {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 1; i <= len(s); i++ {
			prefix := s[:i]
			if isPalindrome(prefix) {
				cur = append(cur, prefix)
				dfs(s[i:], cur)
				cur = cur[:len(cur)-1]
			}
		}

	}
	dfs(s, cur)
	return res
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
