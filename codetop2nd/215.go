package codetop2nd

import (
	"container/heap"
	"sort"
)

type myHeap struct {
	sort.IntSlice
}

func (h myHeap) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *myHeap) Push(i any)        { h.IntSlice = append(h.IntSlice, i.(int)) }
func (h *myHeap) Pop() any {
	tmp := h.IntSlice
	top := tmp[len(tmp)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return top
}

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	minH := &myHeap{}
	heap.Init(minH)
	for _, num := range nums {
		heap.Push(minH, num)
		if minH.Len() > k {
			heap.Pop(minH)
		}
	}
	return minH.IntSlice[0]
}
