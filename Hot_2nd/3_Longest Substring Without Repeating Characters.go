package hot2nd

/*
Wrong Answer
340 / 987 testcases passed

Editorial
Input
s =
" "

Use Testcase
Output
0
Expected
1
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, l, n := 0, 0, len(s)-1
	cur := 0
	ma := make(map[byte]int)
	for l < n {
		if ma[s[l]] == 0 {
			ma[s[l]]++
			cur++
		} else {
			ma[s[l]]++
			cur = 1
			for k := range ma {
				delete(ma, k)
			}
		}
		l++
		res = max(cur, res)
	}
	return res
}

/*
Wrong Answer
407 / 987 testcases passed

Editorial
Input
s =
"dvdf"

Use Testcase
Output
2
Expected
3
*/
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, cur := 0, 0
	ma := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if ma[s[i]] == 0 {
			ma[s[i]]++
			cur++
		} else {
			ma = make(map[byte]int)
			ma[s[i]] = 1
			cur = 1
		}
		if cur > res {
			res = cur
		}
	}
	return res
}

func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, start := 0, 0
	ma := make(map[rune]int)

	for i, c := range s {
		if pos, ok := ma[c]; ok && pos >= start {
			start = pos + 1
		}
		ma[c] = i
		if i-start+1 > res {
			res = i - start + 1
		}
	}
	return res
}
