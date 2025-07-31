package linkedlists

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var prev *ListNode // 用于记录 slow 的前一个节点，以便断开链表

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 此时 slow 指向链表的中间节点 (或者中间偏右一个)
	// prev 指向 slow 的前一个节点
	// 如果链表是偶数个节点，slow 和 fast 会同时到末尾，prev 是前半段的最后一个
	// 如果链表是奇数个节点，fast 会到末尾，slow 是中间节点，prev 是前半段的最后一个

	// 将链表从中点断开
	// prev.Next = nil 使得前半段链表在此处结束
	if prev != nil { // 确保不是只有一个节点的情况 (因为 head != nil || head.Next != nil 已经处理)
		prev.Next = nil
	}

	// 2. 递归排序左右两半
	// leftHalf 是原链表的 head 到 prev
	// rightHalf 是 slow 到原链表末尾
	leftSorted := sortList(head)
	rightSorted := sortList(slow) // slow 现在是右半部分的头节点

	// 3. 合并两个已排序的链表
	return mergeTwoLists_148(leftSorted, rightSorted)
}

// mergeTwoLists 函数用于合并两个排序链表
func mergeTwoLists_148(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{} // 哨兵节点
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}
