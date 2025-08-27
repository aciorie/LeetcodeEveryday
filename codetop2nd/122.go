package codetop2nd

func maxProfit_122(prices []int) int {
	n, res := len(prices), 0
	for i := 0; i < n-1; i++ {
		if prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
