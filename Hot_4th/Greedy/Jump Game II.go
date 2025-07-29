package greedy

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps := 0          // 记录跳跃次数
	currentJumpEnd := 0 // 当前跳跃能够到达的最远边界
	farthestReach := 0  // 从当前位置或之前位置能跳到的最远距离

	for i := 0; i < n; i++ {
		farthestReach = max(farthestReach, i+nums[i])

		if i == currentJumpEnd {
			jumps++
			currentJumpEnd = farthestReach
			if currentJumpEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}
