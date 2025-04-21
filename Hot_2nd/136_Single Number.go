package hot2nd

func singleNumber(nums []int) int {
	res := 0
	if len(nums) == 1 {
		return nums[0]
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for k, v := range freq {
		if v == 1 {
			res = k
		}
	}
	return res
}
