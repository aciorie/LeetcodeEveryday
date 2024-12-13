package hot

import "testing"

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

//this may be a wrong test function
func TestDeleteDuplicates(t *testing.T) {
	inputs := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head:     createList([]int{1, 2, 3, 3, 4, 4, 5}),
			expected: createList([]int{1, 2, 5}),
		},
		{
			head:     createList([]int{1, 1, 1, 2, 3}),
			expected: createList([]int{2, 3}),
		},
	}
	for _, input := range inputs {
		res := DeleteDuplicates(input.head)
		if !compareLists(res, input.expected) {
			t.Errorf("head:%v,expected:%v,got:%v", input.head, input.expected, res)
		}
	}
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
