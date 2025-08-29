package codetop2nd

func deleteDuplicates_83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	slow := dummy

	for fast := head; fast != nil; fast = fast.Next {
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}
		slow.Next = fast
		slow = slow.Next
	}
	return dummy.Next
}
