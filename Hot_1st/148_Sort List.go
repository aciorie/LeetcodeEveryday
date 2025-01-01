package hot

import "sort"

/*
Given the head of a linked list, return the list after sorting it in ascending order.



Example 1:


Input: head = [4,2,1,3]
Output: [1,2,3,4]
Example 2:


Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is in the range [0, 5 * 104].
-105 <= Node.val <= 105


Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/

//original one,using O(n logn) time and O(n) memory
// Runtime	11ms	Beats	67.89%
// Memory	9.26MB	Beats	55.66%
func SortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	res := make([]int, 0)
	tmp := head
	for tmp != nil {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	sort.Ints(res)

	head.Val = res[0]
	tmp = head
	for i := 1; i < len(res); i++ {
		tmp.Next = &ListNode{Val: res[i]}
		tmp = tmp.Next
	}
	return head
}
