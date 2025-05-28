package greedy

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	end := 0
	for i := 0; i < n; i++ {
		if end >= n-1 || i > end {
			break
		}
		end = max(end, i+nums[i])
	}
	return end >= n-1
}
