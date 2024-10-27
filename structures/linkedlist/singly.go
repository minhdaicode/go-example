package linkedlist

import "fmt"

type Singly[T comparable] struct {
	Head   *Node[T]
	length int
}

func NewSingly[T comparable]() *Singly[T] {
	return &Singly[T]{}
}

func (s *Singly[T]) InsertHead(data T) {
	nn := NewNode(data)
	nn.Next = s.Head
	s.Head = nn
	s.length++
}

func (s *Singly[T]) InsertTail(data T) {
	if s.isEmpty() {
		s.InsertHead(data)
		return
	}

	cur := s.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = NewNode(data)
	s.length++
}

func (s *Singly[T]) Insert(data T, postion int) error {
	if postion < 0 || postion > s.length {
		return ErrPositionOutOfRange
	}

	if s.isEmpty() {
		s.InsertHead(data)
		return nil
	}

	cur := s.Head
	for i := 0; cur != nil && i < postion-1; i++ {
		cur = cur.Next
	}

	nn := NewNode(data)
	nn.Next = cur.Next
	cur.Next = nn
	s.length++

	return nil
}

func (s *Singly[T]) Transverse() {
	for cur := s.Head; cur != nil; cur = cur.Next {
		fmt.Printf("%v -> ", cur.Data)
	}
	fmt.Print("nil")
	fmt.Println()
}

func (s *Singly[T]) Len() int {
	return s.length
}

func (s *Singly[T]) isEmpty() bool {
	return s.Head == nil || s.length == 0
}
