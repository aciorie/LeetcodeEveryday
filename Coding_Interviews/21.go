package codinginterviews

/*
面试题 21. 调整数组顺序使奇数位于偶数前面


题目描述
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。



示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：

0 <= nums.length <= 50000
0 <= nums[i] <= 10000
*/
func exchange(nums []int) []int {
	// i points to the current element, and pointer j points to the next position of the last odd number.
	j := 0
	for i, x := range nums {
		if x&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
