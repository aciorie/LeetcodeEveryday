package review

func subarraySum(nums []int, k int) int {
	/*
		我们要找 Sum(nums[i...j]) = k。
		结合前缀和的定义，这意味着 P[j+1] - P[i] = k。
		稍微变换一下：P[j+1] - k = P[i]。
		这就是说，当我们遍历到数组的当前位置 j，并计算出到 j 的前缀和 P[j+1] 时，
		如果之前存在一个前缀和 P[i] 等于 P[j+1] - k，那么就找到了一个符合条件的子数组 nums[i...j]
	*/
	count, currentSum, prefixSumCounts := 0, 0, make(map[int]int)
	prefixSumCounts[0] = 1

	for _, num := range nums {
		currentSum += num
		if _, exist := prefixSumCounts[currentSum-k]; exist {
			count += prefixSumCounts[currentSum-k]
		}
		prefixSumCounts[currentSum]++
	}

	return count
}
