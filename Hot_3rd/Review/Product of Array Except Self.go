package review

func productExceptSelf(nums []int) []int {
	n := len(nums)
	forward, backward, res := make([]int, n), make([]int, n), make([]int, n)
	forward[0], backward[n-1] = 1, 1

	for i := 1; i < n-1; i++ {
		forward[i] = forward[i-1] * nums[i]
	}

	for i := n - 2; i > 0; i-- {
		backward[i] = backward[i+1] * nums[i]
	}

	for i := 0; i < n-1; i++ {
		res[i] = forward[i] * backward[i]
	}
	return res
}
