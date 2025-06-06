package misc

func majorityElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res, maxOccur := 0, 0
	for k, v := range freq {
		if v > maxOccur {
			res = k
			maxOccur = v
		}
	}
	return res
}
