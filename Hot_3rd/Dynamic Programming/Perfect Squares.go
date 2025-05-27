package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			square := j * j
			dp[i] = min(dp[i], dp[i-square]+1)
		}
	}
	return dp[n]
}
