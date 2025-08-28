package codetop2nd

type NodeWithIndex struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*NodeWithIndex{{node: root, index: 1}}
	res := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		leftindex, rightindex := queue[0].index, queue[len(queue)-1].index
		curMax := rightindex - leftindex + 1
		res = max(res, curMax)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.node.Left != nil {
				queue = append(queue, &NodeWithIndex{node: cur.node.Left, index: cur.index * 2})
			}
			if cur.node.Right != nil {
				queue = append(queue, &NodeWithIndex{node: cur.node.Right, index: cur.index*2 + 1})
			}
		}
	}

	return res
}
