package dynamicprogramming

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

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
