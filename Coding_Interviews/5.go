package codinginterviews

import "strings"

/*
面试题 05. 替换空格


题目描述
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



示例 1：

输入：s = "We are happy."
输出："We%20are%20happy."


限制：

0 <= s 的长度 <= 10000
*/

// bulit-in api
func solution5_1(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

// implement it oneself
func solution5_2(s string) string {
	res := strings.Builder{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()
}
