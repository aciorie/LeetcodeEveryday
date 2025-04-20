package codinginterviews

import "strconv"

/*
面试题 46. 把数字翻译成字符串


题目描述
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。



示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"


提示：

0 <= num < 231
*/
func translateNum(num int) int {
	s := strconv.Itoa(num)
	return translateString(s, 0)
}

func translateString(s string, index int) int {
	if index == len(s) {
		return 1
	}
	count := 0

	// Translate a digit
	count += translateString(s, index+1)

	// Try to translate two digits
	if index < len(s)-1 {
		num, _ := strconv.Atoi(s[index : index+2])
		if num >= 10 && num <= 25 {
			count += translateString(s, index+2)
		}
	}

	return count
}
