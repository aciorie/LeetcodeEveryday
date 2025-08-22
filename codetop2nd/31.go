package codetop2nd

func nextPermutation(nums []int) {
	n := len(nums)
	// Step 1: Find the first decreasing element from the end
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 { // If the entire array is not in descending order
		// Step 2: Find the element just larger than nums[i]
		j := n - 1
		for j >= i && nums[i] < nums[j] {
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
