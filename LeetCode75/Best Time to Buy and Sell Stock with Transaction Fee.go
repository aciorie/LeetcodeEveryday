package leetcode75

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	prev_hold, prev_not_hold := -prices[0]-fee, 0
	for i := 1; i < n; i++ {
		cur_hold := max(prev_hold, prev_not_hold-prices[i]-fee)
		cur_not_hold := max(prev_not_hold, prev_hold+prices[i])

		prev_hold = cur_hold
		prev_not_hold = cur_not_hold
	}
	return prev_not_hold
}
