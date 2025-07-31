package linkedlists

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	dummy := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return dummy
}
