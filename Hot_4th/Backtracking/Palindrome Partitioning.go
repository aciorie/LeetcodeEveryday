package backtracking

func partition(s string) [][]string {
	var res [][]string
	var cur []string

	var dfs_131 func(start int)
	dfs_131 = func(start int) {
		if start == len(s) {
			temp := make([]string, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		for i := start; i < len(s); i++ {
			sub := s[start : i+1] // 尝试切分 s[start...i] 作为当前子串

			if isPalindrome(sub) { // 如果子串是回文
				cur = append(cur, sub)
				dfs_131(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}

	dfs_131(0)
	return res
}

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
