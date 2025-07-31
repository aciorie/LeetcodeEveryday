package linkedlists

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	nextPart := reverseList_234(slow.Next)
	for nextPart != nil {
		if head.Val != nextPart.Val {
			return false
		}
		head = head.Next
		nextPart = nextPart.Next
	}
	return true
}

func reverseList_234(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
