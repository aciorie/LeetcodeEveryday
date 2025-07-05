package codetop

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int)
	for i, num := range nums {
		if index, exist := ma[target-num]; exist {
			res[0], res[1] = index, i
			return res
		}
		ma[num] = i
	}
	return res
}
