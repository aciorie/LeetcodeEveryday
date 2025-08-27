package codetop2nd

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondHalf := reverseList_234(slow.Next)
	slow.Next = nil

	firstHalf := head
	res := true
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			res = false
			break
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}
	return res
}

func reverseList_234(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
