package hot

import "math"

/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.



Example 1:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0


Constraints:

1 <= coins.length <= 12
1 <= coins[i] <= 231 - 1
0 <= amount <= 104
*/

//failed
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	res := 0
	for i := len(coins) - 1; i >= 0; i-- {
		if amount == 0 {
			return res
		}

		if coins[i] > amount {
			continue
		} else {
			res += amount / coins[i]
			amount -= amount % coins[i]
		}
	}

	return -1
}

//TLE
func CoinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	res := math.MaxInt32

	var helper func(int, int)
	helper = func(remaining, count int) {
		if remaining < 0 {
			return
		}
		if remaining == 0 {
			res = min(res, count)
			return
		}
		for _, coin := range coins {
			helper(remaining-coin, count+1)
		}
	}
	helper(amount, 0)

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func CoinChange3(coins []int, amount int) int {
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
