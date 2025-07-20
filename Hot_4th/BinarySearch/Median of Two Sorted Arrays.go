package binarysearch

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}
	low, high, halfLen := 0, m, (m+n+1)/2

	for low <= high {
		i := low + (high-low)/2
		j := halfLen - i

		// 定义 L1, R1, L2, R2
		// L1: nums1 左半部分的最大值
		// R1: nums1 右半部分的最小值
		// L2: nums2 左半部分的最大值
		// R2: nums2 右半部分的最小值

		// 使用 math.MinInt64 和 math.MaxInt64 来处理边界情况
		// 例如，如果 i=0，说明 nums1 左边没有元素，L1 视为负无穷大
		L1 := math.MinInt64
		if i != 0 {
			L1 = nums1[i-1]
		}
		R1 := math.MaxInt64
		if i != m {
			R1 = nums1[i]
		}

		L2 := math.MinInt64
		if j != 0 {
			L2 = nums2[j-1]
		}
		R2 := math.MaxInt64
		if j != n {
			R2 = nums2[j]
		}

		if L1 <= R2 && L2 <= R1 {
			// 找到正确的分割点
			// 合并数组的总长度
			totalLen := m + n

			if totalLen%2 == 1 {
				// 如果总长度是奇数，中位数是左半部分的最大值
				return float64(max(L1, L2))
			} else {
				// 如果总长度是偶数，中位数是左半部分的最大值和右半部分的最小值的平均值
				return float64(max(L1, L2)+min(R1, R2)) / 2.0
			}
		} else if L1 > R2 {
			// nums1 的左半部分太大，需要将 i 减小，在 nums1 的左侧继续查找
			high = i - 1
		} else { // L2 > R1
			// nums2 的左半部分太大（或者说 nums1 的右半部分太小），需要将 i 增大，在 nums1 的右侧继续查找
			low = i + 1
		}
	}
	return 0.0
}
