package hashing

func subarraySum(nums []int, k int) int {
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