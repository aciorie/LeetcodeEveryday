package codinginterviews

/*
在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。

Input:
{2, 3, 1, 0, 2, 5}

Output:
2/3
*/

// Hash map
func solution3_1(arr []int) int {
	occur := make(map[int]bool)
	for i := 0; ; i++ {
		if occur[arr[i]] {
			return arr[i]
		}
		occur[arr[i]] = true
	}
}

// Sort
func solution3_2(arr []int) int {
	res := quicksort(arr)
	for i := 0; i < len(res)-1; i++ {
		if res[i] == res[i+1] {
			return res[i]
		}
	}
	return -1
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot, l, r := arr[0], []int{}, []int{}
	for _, v := range arr {
		if v < pivot {
			l = append(l, v)
		} else {
			r = append(r, v)
		}
	}
	return append(append(quicksort(l), pivot), quicksort(r)...)
}

// Swap in place
func solution3_3(nums []int) int {
	// The loop will continue until the return statement is triggered.
	for i := 0; ; i++ {
		// If the value of nums[i] is not equal to i, then the element is not in the correct position.
		for nums[i] != i {
			// The value of the current element
			j := nums[i]
			// If they are equal, it means that the number j is already in the position it should be in and it is a duplicate.
			if nums[j] == j {
				return j
			}
			// Put nums[i] where it should be
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
