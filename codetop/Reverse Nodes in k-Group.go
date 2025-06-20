package codetop

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	dummy := &ListNode{Next: head} // 哨兵节点，方便处理头部的翻转
	prevGroupEnd := dummy          // 指向当前要翻转的组的前一个节点 (翻转后将是新组的头的前一个节点)

	for {
		// --- 1. 检查是否有至少 k 个节点可以翻转 ---
		count := 0
		// `checkNode` 用来遍历，检查当前 `prevGroupEnd.Next` 开始的链表是否有 k 个节点
		checkNode := prevGroupEnd.Next
		for count < k && checkNode != nil {
			checkNode = checkNode.Next
			count++
		}
		// 如果不足 k 个节点，则停止翻转，直接返回
		if count < k {
			break
		}

		// --- 2. 翻转 k 个节点 ---
		// `currentGroupHead` 是当前要翻转的组的第一个节点
		currentGroupHead := prevGroupEnd.Next
		// `currentNode` 是要移动到 `prev` 前面的节点
		currentNode := currentGroupHead.Next

		// 翻转 k-1 次操作（因为第一个节点已经就位）
		for i := 1; i < k; i++ {
			// 保存 `currentNode` 的下一个节点，因为 `currentNode` 的 `Next` 会被改变
			nextNodeToMove := currentNode.Next

			// 核心翻转逻辑
			currentNode.Next = prevGroupEnd.Next   // `currentNode` 指向原组头 (翻转后新组的头)
			prevGroupEnd.Next = currentNode        // `prevGroupEnd` 指向 `currentNode` (将其插入到新组的头部)
			currentGroupHead.Next = nextNodeToMove // 原组头指向 `nextNodeToMove` (维持原组尾部和后续链接)

			// 移动 `currentNode` 到下一个要处理的节点
			currentNode = nextNodeToMove
		}

		// --- 3. 更新 `prevGroupEnd` 到翻转后的组的末尾 ---
		// 翻转前 `currentGroupHead` 是组的头，翻转后它变成了组的尾
		prevGroupEnd = currentGroupHead
	}
	return dummy.Next // 返回新的链表头
}
