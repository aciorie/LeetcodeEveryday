package main

func lengthOfLIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, num := range dp {
		if num > res {
			res = num
		}
	}
	return res
}
