package misc

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n) // 只需要一个 res 数组作为结果和辅助空间

	// 1. 第一次遍历：计算左侧乘积，res[i] 存储的是 nums[0] * ... * nums[i-1] 的乘积
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 2. 第二次遍历：计算右侧乘积，并将其合并到 res 数组中
	rightProduct := 1
	for i := n - 1; i >= 0; i-- { // 从右向左遍历，包括索引 0
		res[i] = res[i] * rightProduct
		rightProduct *= nums[i]
	}

	return res
}
