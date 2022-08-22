package utils

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
