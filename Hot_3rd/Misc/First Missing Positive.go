package misc

func firstMissingPositive(nums []int) int {
	ma := make(map[int]int, len(nums))
	for _, num := range nums {
		ma[num]++
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := ma[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}
