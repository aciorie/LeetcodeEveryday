package codetop

func merge_forget(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	// p1 指向 nums1 有效部分的末尾
	p1 := m - 1
	// p2 指向 nums2 的末尾
	p2 := n - 1
	// p 指向 nums1 中合并后数组的末尾，也就是要填充的位置
	p := m + n - 1

	// 当 nums2 中还有元素需要处理时，循环继续
    for p2 >= 0 {
        // 如果 nums1 还有元素可以比较，并且 nums1 的当前元素大于 nums2 的当前元素
        if p1 >= 0 && nums1[p1] > nums2[p2] {
            nums1[p] = nums1[p1] // 将 nums1 的元素放到正确位置
            p1--                 // nums1 指针左移
        } else {
            // 否则（nums1 已经处理完，或者 nums2 的元素更大/相等）
            nums1[p] = nums2[p2] // 将 nums2 的元素放到正确位置
            p2--                 // nums2 指针左移
        }
        p-- // 填充位置指针左移
    }
}
