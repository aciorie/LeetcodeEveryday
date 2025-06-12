package review

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}
type MinHeap_23 []int

func (h MinHeap_23) Len() int           { return h.Len() }
func (h MinHeap_23) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap_23) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap_23) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap_23) Pop() any {
	tmp := *h
	top := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return top
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	h := &MinHeap_23{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	// 构建合并后的排序链表
	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)

		tail.Next = minNode // 将取出的最小节点连接到结果链表的末尾
		tail = minNode      // 移动 tail 指针到刚刚添加的新节点

		// 如果这个最小节点所在的原始链表还有后续节点，将其下一个节点推入堆中
		// 作为新的候选者，参与下一轮的最小值比较
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummy.Next
}
