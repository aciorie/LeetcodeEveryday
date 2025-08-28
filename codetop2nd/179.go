package codetop2nd

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
