package linkedlist

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{data, nil}
}

type DoublyNode[T comparable] struct {
	Data T
	Pre  *DoublyNode[T]
	Next *DoublyNode[T]
}

func NewDoublyNode[T comparable](data T) *DoublyNode[T] {
	return &DoublyNode[T]{data, nil, nil}
}
