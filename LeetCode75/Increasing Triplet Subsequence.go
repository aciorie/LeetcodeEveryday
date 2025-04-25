package leetcode75

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= first {
			first = num
		}
		if first < num && num <= second {
			second = num
		}
		if num > second {
			return true
		}
	}
	return false
}
