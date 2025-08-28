package codetop2nd

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	res := nums[0]
	maxValue, minValue := make([]int, n), make([]int, n)
	maxValue[0], minValue[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		cur := nums[i]
		pro1 := maxValue[i-1] * nums[i]
		pro2 := minValue[i-1] * nums[i]

		maxValue[i] = max(cur, max(pro1, pro2))
		minValue[i] = max(cur, min(pro1, pro2))

		res = max(res, maxValue[i])
	}
	return res
}
