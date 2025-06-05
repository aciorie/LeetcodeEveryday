package misc

func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 { // 如果数组为空或只有一个元素，没有下一个排列
		return
	}

	// 1. 从右向左查找第一个“下降”的元素 (找到枢轴 i)
	// 这个 i 应该满足 nums[i] < nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] { // 注意这里是 nums[i] >= nums[i+1]
		i--
	}

	// 如果找到了枢轴 i (即 i >= 0)
	if i >= 0 {
		// 2. 从右向左查找第一个比枢轴大的元素 (找到交换元素 j)
		// 这个 j 应该满足 nums[j] > nums[i]
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] { // 注意这里是 nums[j] <= nums[i]
			j--
		}
		// 3. 交换 nums[i] 和 nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 4. 反转枢轴后面的后缀 (从 i+1 到数组末尾)
	// 无论是否找到枢轴 i (即数组是完全降序的)，都需要反转
	// 如果没找到 i，i 会是 -1，那么需要反转整个数组 (从 0 开始)
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
