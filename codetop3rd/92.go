package codetop3rd

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	start := pre.Next
	cur := start.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
	}
	return dummy.Next
}
