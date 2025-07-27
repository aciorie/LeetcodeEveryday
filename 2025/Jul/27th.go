package jul

func countHillValley(nums []int) int {
	var processedNums []int
	if len(nums) > 0 {
		processedNums = append(processedNums, nums[0])
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				processedNums = append(processedNums, nums[i])
			}
		}
	}

	if len(processedNums) < 3 {
		return 0
	}

	count := 0
	for i := 1; i < len(processedNums)-1; i++ {
		if processedNums[i] > processedNums[i-1] && processedNums[i] > processedNums[i+1] {
			count++
		} else if processedNums[i] < processedNums[i-1] && processedNums[i] < processedNums[i+1] {
			count++
		}
	}

	return count
}
