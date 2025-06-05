package misc

func maxSubArray(nums []int) int {
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
