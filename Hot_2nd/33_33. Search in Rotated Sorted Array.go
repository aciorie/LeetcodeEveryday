package hot2nd

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	n := len(nums)
	l, r := 0, n-1
	for l <= r { //  for the possibility that len(nums)==1
		m := l + (r-l)/2
		if target == nums[m] {
			return m
		}

		// left side is ordered
		if nums[l] <= nums[m] {
			// if target in the left side
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else { // target not in the left side
				l = m + 1
			}
		} else { // right side is ordered
			// if target in the right side
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				// target not in the left side
				r = m - 1
			}
		}
	}
	return -1
}
