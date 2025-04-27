package leetcode75

import "math"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if n < k || k <= 0 {
		return 0.0
	}

	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k; i < n; i++ {
		currentSum = currentSum - nums[i-k] + nums[i]

		maxSum = int(math.Max(float64(maxSum), float64(currentSum)))
	}

	return float64(maxSum) / float64(k)
}
