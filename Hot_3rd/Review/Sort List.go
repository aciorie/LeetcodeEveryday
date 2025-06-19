package review

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 0
	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	dummy := &ListNode{Next: head}

	for subLength := 1; subLength < n; subLength *= 2 {
		prev := dummy
		curr = dummy.Next

		for curr != nil {
			l1 := curr
			temp := l1
			for i := 1; i < subLength && temp != nil; i++ {
				temp = temp.Next
			}

			if temp == nil {
				prev.Next = l1
				break
			}

			l2 := temp.Next
			temp.Next = nil

			temp = l2
			for i := 1; i < subLength && temp != nil; i++ {
				temp = temp.Next
			}

			var nextUnmergedHead *ListNode
			if temp != nil {
				nextUnmergedHead = temp.Next
				temp.Next = nil
			} else {
				nextUnmergedHead = nil
			}

			mergedHead, mergedTail := mergeTwoLists(l1, l2)

			prev.Next = mergedHead
			prev = mergedTail

			curr = nextUnmergedHead
		}
	}

	return dummy.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
	if l1 == nil {
		return l2, getTail(l2)
	}
	if l2 == nil {
		return l1, getTail(l1)
	}

	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	for tail.Next != nil {
		tail = tail.Next
	}

	return dummy.Next, tail
}

func getTail(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for head.Next != nil {
		head = head.Next
	}
	return head
}
