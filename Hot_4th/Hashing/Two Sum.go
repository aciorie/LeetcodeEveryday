package hashing

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	ma := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, exist := ma[target-num]; exist {
			res[0] = i
			res[1] = index
			return res
		}
		ma[num] = i
	}
	return res
}
