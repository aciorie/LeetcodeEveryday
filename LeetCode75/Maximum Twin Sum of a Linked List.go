package leetcode75

func pairSum(head *ListNode) int {
	secondHalfStart := findMiddle(head)
	reversedSecondHalf := reverseList(secondHalfStart)
	firstHalfPtr := head
	secondHalfPtr := reversedSecondHalf
	res := 0

	for secondHalfPtr != nil {
		currentSum := firstHalfPtr.Val + secondHalfPtr.Val

		if currentSum > res {
			res = currentSum
		}

		firstHalfPtr = firstHalfPtr.Next
		secondHalfPtr = secondHalfPtr.Next
	}
	return res
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}
