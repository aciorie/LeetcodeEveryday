package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// slow, fast := head, head
	// for i := 0; i < n; i++ {
	// 	fast = fast.Next
	// }
	// for fast != nil {
	// 	slow = slow.Next
	// 	fast = fast.Next
	// }
	// wantNext := slow.Next.Next
	// slow.Next = wantNext
	// return head

	dummy := &ListNode{Next: head}
	slow, fast := dummy, head

	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
