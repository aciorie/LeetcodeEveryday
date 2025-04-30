package leetcode75

func findDifference(nums1 []int, nums2 []int) [][]int {
	hash1, hash2 := make(map[int]bool, len(nums1)), make(map[int]bool, len(nums2))
	for _, num := range nums1 {
		hash1[num] = true
	}
	for _, num := range nums2 {
		hash2[num] = true
	}

	var list1Only []int
	var list2Only []int

	for num := range hash1 {
		if _, exists := hash2[num]; !exists {
			list1Only = append(list1Only, num)
		}
	}

	for num := range hash2 {
		if _, exists := hash1[num]; !exists {
			list2Only = append(list2Only, num)
		}
	}

	return [][]int{list1Only, list2Only}
}
