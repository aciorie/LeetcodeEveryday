package codetop2nd

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			if dp[j-coin] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
