package codetop3rd

func quickSort(arr []int) {
	if len(arr) > 1 {
		quickSortRecursive(arr, 0, len(arr)-1)
	}
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortRecursive(arr, low, pivotIndex-1)
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
