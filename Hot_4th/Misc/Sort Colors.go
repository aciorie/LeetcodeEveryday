package misc

func sortColors(nums []int) {
	zeros, ones, twos := 0, 0, len(nums)-1
	for ones <= twos {
		switch nums[ones] {
		case 0:
			nums[zeros], nums[ones] = nums[ones], nums[zeros]
			zeros++
			ones++
		case 1:
			ones++
		case 2:
			nums[twos], nums[ones] = nums[ones], nums[twos]
			twos--
		}
	}
}
