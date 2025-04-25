package leetcode75

func productExceptSelf(nums []int) []int {
	n := len(nums)
	forward, backward, res := make([]int, n), make([]int, n), make([]int, n)

	forward[0], backward[n-1] = 1, 1

	for i := 0; i < n-1; i++ {
		forward[i+1] = forward[i] * nums[i]
	}

	for i := n - 1; i > 0; i-- {
		backward[i-1] = backward[i] * nums[i]
	}

	for i := 0; i < n; i++ {
		res[i] = forward[i] * backward[i]
	}

	return res
}
