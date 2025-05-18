package leetcode75

func singleNumber(nums []int) int{
	if len(nums) == 1 {
		return nums[0]
	}
	ma := make(map[int]int)
	for _, num := range nums {
		ma[num]++
	}
	for k, v := range ma {
		if v == 1 {
			return k
		}
	}
	return 0
}