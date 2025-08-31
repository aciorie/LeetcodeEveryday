package codetop3rd

import (
	"container/heap"
	"sort"
)

type myHeap struct{ sort.IntSlice }

func (h myHeap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func (h *myHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *myHeap) Pop() any {
	tmp := h.IntSlice
	h.IntSlice = tmp[:len(tmp)-1]
	return tmp[len(tmp)-1]
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	minH := &myHeap{}
	for _, num := range nums {
		heap.Push(minH, num)
		if minH.Len() > k {
			heap.Pop(minH)
		}
	}
	return minH.IntSlice[0]
}
