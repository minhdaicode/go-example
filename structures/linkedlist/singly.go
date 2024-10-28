package linkedlist

import (
	"fmt"
)

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

func (s *Singly[T]) Insert(data T, position int) error {
	if position < 0 || position > s.length {
		return ErrPositionOutOfRange
	}

	if s.isEmpty() {
		s.InsertHead(data)
		return nil
	}

	cur := s.Head
	for i := 0; cur != nil && i < position-1; i++ {
		cur = cur.Next
	}

	nn := NewNode(data)
	nn.Next = cur.Next
	cur.Next = nn
	s.length++
	return nil
}

func (s *Singly[T]) RemoveHead() error {
	if s.isEmpty() {
		return ErrListEmpty
	}

	cur := s.Head
	s.Head = cur.Next
	s.length--
	return nil
}

func (s *Singly[T]) RemoveTail() error {
	if s.isEmpty() {
		return ErrListEmpty
	}

	if s.Head.Next == nil {
		return s.RemoveHead()
	}

	cur := s.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	s.length--
	return nil
}

func (s *Singly[T]) Remove(position int) error {
	if position < 0 || position > s.length {
		return ErrPositionOutOfRange
	}

	if s.isEmpty() {
		return ErrListEmpty
	}

	if position == 0 {
		return s.RemoveHead()
	}

	cur := s.Head
	for i := 0; cur != nil && i < position-1; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	s.length--

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
