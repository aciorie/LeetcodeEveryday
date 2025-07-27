package dynamicprogramming

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := nums[0]
	maxValue, minValue := make([]int, n), make([]int, n)
	maxValue[0], minValue[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		prod1, prod2 := maxValue[i-1]*cur, minValue[i-1]*cur

		maxValue[i] = max(cur, max(prod1, prod2))
		minValue[i] = min(cur, min(prod1, prod2))

		res = max(res, maxValue[i])
	}

	return res
}
