package main

func moveZeroes(nums []int) {
	l, r := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
