package review

import "sort"

// false one
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sort.Ints(nums)
	leftSum, rightSum := nums[0], nums[n-1]
	l, r := 1, n-1
	for l <= r {
		if leftSum == rightSum {
			return true
		} else if leftSum < rightSum {
			leftSum += nums[l]
			l++
		} else {
			rightSum += nums[r]
			r--
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 != 0 {
		return false
	}

	// 目标和：每个子集需要达到的和
	target := totalSum / 2

	// dp 数组：dp[i] 表示是否能从 nums 中组合成和为 i
	// 数组大小为 target + 1，因为索引 0 到 target
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}
