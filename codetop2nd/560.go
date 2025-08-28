package codetop2nd

func subarraySum(nums []int, k int) int {
	res, cur, prefix := 0, 0, make(map[int]int)
	prefix[0] = 1

	for _, num := range nums {
		cur += num
		if _, exist := prefix[cur-k]; exist {
			res += prefix[cur-k]
		}
		prefix[cur]++
	}
	return res
}
