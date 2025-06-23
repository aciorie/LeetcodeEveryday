package review

func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	// Step 1: Find the first decreasing element from the end
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 { // If the entire array is not in descending order
		// Step 2: Find the element just larger than nums[i]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// Step 3: Swap elements at i and j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 4: Reverse the sequence from i + 1 to the end
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	end := len(nums) - 1
	for end > start {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
