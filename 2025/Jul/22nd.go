package jul

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 0 {
		return 0
	}

	curSum, maxSum := 0, 0
	appeared := make(map[int]bool)
	l := 0

	for r := 0; r < n; r++ {
		for appeared[nums[r]] {
			appeared[nums[l]] = false
			curSum -= nums[l]
			l++
		}
		appeared[nums[r]] = true
		curSum += nums[r]

		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
