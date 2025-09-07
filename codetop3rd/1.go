package codetop3rd

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int)
	for k, v := range nums {
		if index, exists := ma[target-v]; exists {
			res[0] = index
			res[1] = k
			return res
		}
		ma[v] = k
	}
	return res
}
