package helpers

type LinkedListNode struct {
	data interface{}
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	size int
}
