package greedy

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	end := 0
	for i := 0; i < n; i++ {
		if i > end {
			return false
		}
		end = max(end, i+nums[i])
		if end >= n-1 {
			return true
		}
	}
	return false
}
