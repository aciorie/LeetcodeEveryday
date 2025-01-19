package hot

/*
Given the head of a singly linked list, return true if it is a palindrome or false otherwise.



Example 1:
Input: head = [1,2,2,1]
Output: true

Example 2:
Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9


Follow up: Could you do it in O(n) time and O(1) space?
*/

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for l, r := 0, len(arr)-1; l < r; {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
