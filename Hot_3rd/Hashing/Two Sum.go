package hashing

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int)
	for k, v := range nums {
		if i, exists := ma[target-v]; exists {
			res[0], res[1] = i, k
			return res
		}
		ma[v] = k
	}
	return res
}
