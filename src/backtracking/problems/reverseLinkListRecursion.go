package backtracking

import "interviewBit/src/utils"

/*
	Problem : https://www.interviewbit.com/problems/reverse-link-list-recursion/
*/

func solve(prev *utils.ListNode, curr *utils.ListNode) *utils.ListNode {
	if curr == nil {
		return prev
	}

	currNext := curr.Next
	curr.Next = prev
	return solve(curr, currNext)
}

// T(n) : O(n), S(n) : O(n)
func ReverseLinkedList(A *utils.ListNode) *utils.ListNode {
	return solve(nil, A)
}
