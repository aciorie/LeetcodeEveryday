package codetop4th

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}
