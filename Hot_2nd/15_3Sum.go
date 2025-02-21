package hot2nd

func threeSum(nums []int) [][]int {
	nums = QuickSort(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	l, m, r := []int{}, []int{}, []int{}
	for _, x := range arr {
		if x < pivot {
			l = append(l, x)
		} else if x == pivot {
			m = append(m, x)
		} else {
			r = append(r, x)
		}
	}
	return append(append(QuickSort(l), m...), QuickSort(r)...)
}
