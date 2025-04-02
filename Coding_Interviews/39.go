package codinginterviews

/*
面试题 39. 数组中出现次数超过一半的数字


题目描述
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。



你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2


限制：

1 <= 数组长度 <= 50000
*/
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res, maxOccur := 0, 0
	for k, v := range freq {
		if v > maxOccur {
			res = k
			maxOccur = v
		}
	}
	return res
}
