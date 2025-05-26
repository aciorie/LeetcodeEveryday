package main

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxF := make([]int, n)
	minF := make([]int, n)
	res := nums[0]

	maxF[0] = nums[0]
	minF[0] = nums[0]

	for i := 1; i < n; i++ {
		current := nums[i]
		prod1 := maxF[i-1] * current
		prod2 := minF[i-1] * current

		maxF[i] = max(current, max(prod1, prod2))
		minF[i] = min(current, min(prod1, prod2))

		res = max(res, maxF[i])
	}

	return res
}
