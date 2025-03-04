package hot2nd

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]ListNode, 0)
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head
	tmp := k
	for tmp > 0 {
		cur = cur.Next
		prev = prev.Next
		stack = append(stack, *cur)
		tmp--
	}

	cur2 := cur
	for k > 0 {
		cur.Next = prev
		prev.Next = &stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
	}
	return cur2
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		// Check if there are at least k nodes left in the list
		count := 0
		current := prevGroupEnd.Next
		for count < k && current != nil {
			current = current.Next
			count++
		}
		if count < k {
			break // Not enough nodes to reverse, exit the loop
		}

		// Reverse k nodes
		prev := prevGroupEnd.Next
		curr := prev.Next
		for i := 1; i < k; i++ {
			prev.Next = curr.Next
			curr.Next = prevGroupEnd.Next
			prevGroupEnd.Next = curr
			curr = prev.Next
		}

		// Move prevGroupEnd to the end of the reversed group
		prevGroupEnd = prev
	}

	return dummy.Next
}
