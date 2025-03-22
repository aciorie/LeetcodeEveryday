package hot2nd

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	n, curMax, res := len(nums), nums[0], nums[0]
	for i := 1; i < n; i++ {
		curMax = max(nums[i], curMax+nums[i])
		res = max(res, curMax)
	}
	return res
}
