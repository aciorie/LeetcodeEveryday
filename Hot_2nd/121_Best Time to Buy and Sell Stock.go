package hot2nd

func maxProfit(prices []int) int {
	// The maximum profit stored up to each day
	dp := make([]int, len(prices))
	dp[0] = 0
	// Lowest price so far
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dp[len(dp)-1]
}

func maxProfit2(prices []int) int {
	maxProfi, minPrice := 0, prices[0]
	for _, price := range prices {
		maxProfi = max(maxProfi, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxProfi
}
