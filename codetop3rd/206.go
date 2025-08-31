package codetop3rd

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return dummy
}
