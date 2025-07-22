package codetop

func maxProfit(prices []int) int {
	n := len(prices)
	maxProfit, minPrice := 0, prices[0]
	for i := 1; i < n; i++ {
		maxProfit = max(maxProfit, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfit
}
