package hot3rd

// false one
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if len(s) == 1 {
		return [][]string{{string(s[0])}}
	}

	var res [][]string
	var cur []string
	var backtrack func(s string, cur []string)
	backtrack = func(s string, cur []string) {
		if ispalindrome(cur) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(s); i++ {
			cur = append(cur, string(s[i]))
			backtrack(s, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(s, cur)
	return res
}

func ispalindrome(cur []string) bool {
	if len(cur) < 2 {
		return true
	}
	l, r := 0, len(cur)
	for l < r {
		if cur[l] == cur[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
