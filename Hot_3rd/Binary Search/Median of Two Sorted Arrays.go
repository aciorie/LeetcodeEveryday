package hot3rd

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	m := len(nums1)
	if m%2 == 1 {
		return float64(nums1[m/2])
	} else {
		return float64(nums1[m/2]+nums1[m/2-1]) / 2.0
	}
}
