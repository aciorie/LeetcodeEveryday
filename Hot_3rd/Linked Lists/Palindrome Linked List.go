package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	value := make([]int, 0)
	for head != nil {
		value = append(value, head.Val)
		head = head.Next
	}

	for l, r := 0, len(value)-1; l < r; {
		if value[l] != value[r] {
			return false
		}
		l++
		r--
	}
	return true
}
