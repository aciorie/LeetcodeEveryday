package codinginterviews

import "math"

/*
面试题 42. 连续子数组的最大和


题目描述
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
*/

// Wrong one
func maxSubArray_1(nums []int) int{
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	return dp[n-1]
}

// fixed bugs
func maxSubArray_2(nums []int) int {
	n := len(nums)
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

// A better solution
func maxSubArray_3(nums []int) int{
	res, cur := math.MinInt32, 0
    for _, num := range nums {
        cur = max(cur, 0) + num
        res = max(res, cur)
    }
    return res
}