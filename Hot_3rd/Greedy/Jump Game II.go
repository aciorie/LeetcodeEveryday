package greedy

func jump(nums []int) int {
	n, maxReach := len(nums), 0
	end, steps := 0, 0

	for i := 0; i < n-1; i++ {
		maxReach = max(maxReach, i+nums[i])
		if i == end {
			end = maxReach
			steps++
		}
	}

	return steps
}
