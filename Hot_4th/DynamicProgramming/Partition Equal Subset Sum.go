package dynamicprogramming

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	if totalSum%2 == 1 {
		return false
	}

	targetSum := totalSum / 2
	dp := make([]bool, targetSum+1)
	dp[0] = true

	for _, num := range nums {
		for j := targetSum; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[targetSum]
}
