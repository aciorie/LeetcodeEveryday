package codetop2nd

func maxProfit(prices []int) int {
	n := len(prices)
	minp, maxp := prices[0], 0
	for i := 1; i < n; i++ {
		maxp = max(maxp, prices[i]-minp)
		minp = min(minp, prices[i])
	}
	return maxp
}
