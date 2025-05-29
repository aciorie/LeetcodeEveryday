package heap

import (
	"container/heap"
	"sort"
)

type MyHeap struct{ sort.IntSlice }

func (h MyHeap) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *MyHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MyHeap) Pop() any {
	tmp := h.IntSlice
	top := tmp[len(tmp)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return top
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	minH := &MyHeap{}
	for _, num := range nums {
		heap.Push(minH, num)
		if minH.Len() > k {
			heap.Pop(minH)
		}
	}

	return minH.IntSlice[0]
}
