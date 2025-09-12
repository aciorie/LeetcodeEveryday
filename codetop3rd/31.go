package codetop3rd

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	i := n - 2
	for i >= 0 && nums[i] >= nums[i-1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse_31(nums, i+1)
}

func reverse_31(nums []int, start int) {
	end := len(nums) - 1
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
