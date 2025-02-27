package hot2nd

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, slow, fast := dummy, head, head
	for fast != nil {
		if n <= 0 {
			pre = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	pre.Next = slow.Next
	return dummy.Next
}
