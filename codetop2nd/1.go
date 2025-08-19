package codetop2nd

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int)
	for i, num := range nums {
		if index, exists := ma[target-num]; exists {
			res[0] = index
			res[1] = i
			return res
		}
		ma[num] = i
	}
	return res
}
