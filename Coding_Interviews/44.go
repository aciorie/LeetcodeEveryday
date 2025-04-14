package codinginterviews

import "strconv"

/*面试题 44. 数字序列中某一位的数字


题目描述
数字以123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。



示例 1：

输入：n = 3
输出：3
示例 2：

输入：n = 11
输出：0


限制：

0 <= n < 2^31
*/

/*
Break the problem down into the following steps:
	First, determine the number of digits in the number.
	Then, identify the starting number for the target digits. For example, for two-digit numbers,
		the starting number is 10; for three-digit numbers, the starting number is 100.
	Next, determine the specific target number.
	Finally, identify the target position within the target number.
*/
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	digits, count, start := 1, 9, 1

	// Determine the number of digits
	for n > count*digits {
		n -= count * digits
		digits++
		count *= 10
		start *= 10 // Update the starting number
	}

	// Determine the target number
	number := start + (n-1)/digits

	// Determine the target position in the target number
	index := (n - 1) % digits
	s := strconv.Itoa(number)                  // Convert a number to a string
	digit, _ := strconv.Atoi(string(s[index])) // Take out the target bit and convert it into a number

	return digit
}
