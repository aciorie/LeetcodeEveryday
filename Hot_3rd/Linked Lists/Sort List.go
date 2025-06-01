package main

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := make([]int, 0)
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	sort.Ints(res)

	dummy := &ListNode{}
	tail := dummy
	for i := 0; i < len(res); i++ {
		newNode := &ListNode{Val: res[i]}
		tail.Next = newNode
		tail = newNode
	}
	return dummy.Next
}
