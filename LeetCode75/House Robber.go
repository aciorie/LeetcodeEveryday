package leetcode75

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, n)
	dp[0], dp[1], dp[2] = nums[0], nums[1], nums[0]+nums[2]

	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}

	return max(dp[n-1], dp[n-2])
}
