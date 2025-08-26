package codetop2nd

func searchRange(nums []int, target int) []int {
	leftBound := findLeftBound(nums, target)
	rightBound := findRightBound(nums, target)

	return []int{leftBound, rightBound}
}

func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	leftmost := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			leftmost = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return leftmost
}

func findRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightmost := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			rightmost = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return rightmost
}
