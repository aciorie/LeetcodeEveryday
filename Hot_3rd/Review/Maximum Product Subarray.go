package review

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := nums[0]
	maxValue, minValue := make([]int, n), make([]int, n)
	maxValue[0], minValue[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		current := nums[i]
		prod1 := maxValue[i-1] * current
		prod2 := minValue[i-1] * current

		maxValue[i] = max(current, max(prod1, prod2))
		minValue[i] = min(current, min(prod1, prod2))

		res = max(res, maxValue[i])
	}

	return res
}
