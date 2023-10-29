package utils

type ListNode struct {
	Value int
	Next  *ListNode
}

func ListNode_new(val int) *ListNode {
	var node *ListNode = new(ListNode)
	node.Value = val
	node.Next = nil
	return node
}

// T(n) : O(n), S(n) : O(1)
func CompareLinkedLists(A *ListNode, B *ListNode) bool {

	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		return false
	}

	for A != nil && B != nil {
		if A.Value != B.Value {
			return false
		}
		A = A.Next
		B = B.Next
	}

	return A == nil && B == nil
}

// T(n) : O(n), S(n) : O(n)
func GenerateLinkedList(A []int) *ListNode {

	n := len(A)
	if n == 0 {
		return nil
	}
	var head *ListNode = ListNode_new(A[0])
	var currNode *ListNode = head
	for i := 1; i < n; i++ {
		currNode.Next = ListNode_new(A[i])
		currNode = currNode.Next
	}

	return head
}

type LinkedListNode struct {
	Data interface{}
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	size int
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) InsertAtHead(v int) (*LinkedList, bool) {
	lLN := new(LinkedListNode)
	lLN.Data = v
	lLN.Next = nil

	if l == nil || l.size == 0 {
		l = new(LinkedList)
		l.Head = lLN
	} else {
		l.Head.Next = lLN
	}
	l.size += 1

	return l, true
}

// func (l *LinkedList) InsertAtTail(v int) (*LinkedList, bool) {
// }

// func NextNode(n *LinkedListNode) (*LinkedListNode, bool) {
// 	if n.Next == nil {
// 		return nil, false
// 	}
// 	return n.Next, true
// }

// func (l *LinkedList) IsNodePresent(n *LinkedListNode) bool {
// 	for _, ok := NextNode(n); if ok {

// 	}
// 	return false
// }
