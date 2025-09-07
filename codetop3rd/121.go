package codetop3rd

func maxProfit(prices []int) int {
	n := len(prices)
	maxprofit, minprice := 0, prices[0]
	for i := 1; i < n; i++ {
		maxprofit = max(maxprofit, prices[i]-minprice)
		minprice = min(minprice, prices[i])
	}
	return maxprofit
}
