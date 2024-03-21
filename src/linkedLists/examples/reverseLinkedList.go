package linkedListExamples

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/reverse-linked-list/
*/

func reverse(prev *utils.ListNode, curr *utils.ListNode, currNext *utils.ListNode) *utils.ListNode {

	if currNext == nil {
		return curr
	}

	var temp *utils.ListNode = currNext.Next
	currNext.Next = curr
	curr.Next = prev

	return reverse(curr, currNext, temp)
}

// T(n) : O(n), S(n) : O(n)
func ReverseList(A *utils.ListNode) *utils.ListNode {
	return reverse(nil, A, A.Next)
}
