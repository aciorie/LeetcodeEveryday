package linkedlists

type Node_R struct {
	Val    int
	Next   *Node_R
	Random *Node_R
}

func copyRandomList(head *Node_R) *Node_R {
	if head == nil {
		return nil
	}

	oldToNewMap := make(map[*Node_R]*Node_R)
	cur := head
	for cur != nil {
		oldToNewMap[cur] = &Node_R{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := oldToNewMap[cur]
		if cur.Next != nil {
			newNode.Next = oldToNewMap[cur.Next]
		}
		if cur.Random != nil {
			newNode.Random = oldToNewMap[cur.Random]
		}

		cur = cur.Next
	}
	return oldToNewMap[head]
}
