package review

func partition(s string) [][]string {
	var res [][]string
	var currentPartition []string

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(currentPartition))
			copy(temp, currentPartition)
			res = append(res, temp)
			return
		}

		// 遍历所有可能的结束点 i，切分出子串
		for i := start; i < len(s); i++ {
			sub := s[start : i+1] // 尝试切分 s[start...i] 作为当前子串

			if isPalindrome(sub) { // 如果子串是回文
				currentPartition = append(currentPartition, sub)
				backtrack(i + 1)
				currentPartition = currentPartition[:len(currentPartition)-1]
			}
		}
	}

	backtrack(0)

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
