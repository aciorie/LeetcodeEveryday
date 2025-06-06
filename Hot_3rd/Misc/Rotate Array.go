package misc

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	k = k % n

	if k == 0 {
		return
	}

	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}
