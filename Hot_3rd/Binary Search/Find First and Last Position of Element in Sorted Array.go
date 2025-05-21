package main

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)

	if n == 0 {
		return res
	}

	left, right := 0, n-1
	foundIndex := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			foundIndex = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if foundIndex == -1 {
		return res
	}

	res[0] = foundIndex

	left, right = 0, n-1
	lastIndex := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			lastIndex = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	res[1] = lastIndex

	return res
}
