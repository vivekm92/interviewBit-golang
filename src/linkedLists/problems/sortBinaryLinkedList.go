package linkedList

import (
	"interviewBit/src/utils"
)

/*
	problem : https://www.interviewbit.com/problems/sort-binary-linked-list/
*/

// T(n) : O(n), S(n), O(n)
func SortBinaryLinkedList(A *utils.ListNode) *utils.ListNode {

	if A == nil {
		return nil
	}

	cnt0, cnt1 := 0, 0
	var head *utils.ListNode = utils.ListNode_new(-1)
	var curr *utils.ListNode = head
	for A != nil {
		if A.Value == 0 {
			cnt0++
		} else {
			cnt1++
		}
		A = A.Next
	}

	for cnt0 > 0 {
		curr.Next = utils.ListNode_new(0)
		curr = curr.Next
		cnt0--
	}

	for cnt1 > 0 {
		curr.Next = utils.ListNode_new(1)
		curr = curr.Next
		cnt1--
	}

	return head.Next
}
