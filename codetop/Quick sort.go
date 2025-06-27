package codetop

func quicksort(nums []int) {
	if len(nums) < 2 {
		return
	}
	quicksortInternal(nums, 0, len(nums)-1)
}

func quicksortInternal(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quicksortInternal(nums, low, pivotIndex-1)
		quicksortInternal(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1 // i 是分区点，它的左边（包括它自己）存放小于或等于基准的元素

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
