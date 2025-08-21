package codetop2nd

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	wantNext := slow.Next.Next
	slow.Next = wantNext
	return head
}
