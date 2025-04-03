package hot2nd

func sortColors1(nums []int) {
	freq := make(map[int]int, 3)
	n := len(nums)
	for _, num := range nums {
		freq[num]++
	}
	for i := 0; i < n; i++ {
		if i < freq[0] {
			nums[i] = 0
		} else if i < freq[0]+freq[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func sortColors2(nums []int) {
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
			nums[ones], nums[twos] = nums[twos], nums[ones]
			// To check the new exchanged num, dont ones++
			// ones++
			twos--
		}
	}
}
