package misc

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	if k == 0 {
		return
	}
	reverse_189(nums, 0, n-1)
	reverse_189(nums, 0, k-1)
	reverse_189(nums, k, n-1)
}

func reverse_189(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
