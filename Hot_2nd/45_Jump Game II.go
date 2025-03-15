package hot2nd

import "math"

// dp
func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// If it is possible to jump from j to i
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

// greedy
func jump2(nums []int) int {
	n, maxReach := len(nums), 0
	// The boundary of the jump range of the current step, the number of jumps
	end, steps := 0, 0

	for i := 0; i < n-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		// If the boundary of the current jump range is reached
		if i == end {
			// Enter the next step jump range
			end = maxReach
			steps++
		}
	}

	return steps
}
