package dynamicprogramming

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n]
}
