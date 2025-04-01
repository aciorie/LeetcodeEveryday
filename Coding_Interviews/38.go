package codinginterviews

import "sort"

/*
面试题 38. 字符串的排列


题目描述
输入一个字符串，打印出该字符串中字符的所有排列。



你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。



示例:

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]


限制：

1 <= s 的长度 <= 8
*/
func permutation(s string) []string {
	res, cur, used := []string{}, "", make([]bool, len(s))
	chars := []rune(s)

	sort.Slice(chars, func(i, j int) bool {
		return s[i] < s[j]
	})

	var dfs func()
	dfs = func() {
		if len(cur) == len(s) {
			res = append(res, cur)
		}
		for i := 0; i < len(s); i++ {
			if used[i] {
				continue
			}
			if i > 0 && chars[i] == chars[i-1] && !used[i-1] {
				continue
			}

			cur += string(chars[i])
			used[i] = true
			dfs()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
		// used set true

		// used set false
	}
	dfs()
	return res
}
