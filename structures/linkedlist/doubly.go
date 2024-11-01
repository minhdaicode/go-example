package linkedlist

import (
	"fmt"
)

type Doubly[T comparable] struct {
	Head   *DoublyNode[T]
	Tail   *DoublyNode[T]
	length int
}

func NewDoubly[T comparable]() *Doubly[T] {
	return &Doubly[T]{}
}

func (d *Doubly[T]) InsertHead(data T) {
	nn := NewDoublyNode(data)

	switch {
	case d.isEmpty():
		d.Head = nn
		d.length++
		return

	case d.Head.Next == nil:
		d.Tail = d.Head
		d.Head = nn
		d.Head.Next = d.Tail
		d.Tail.Pre = d.Head
		d.length++
		return

	default:
		nn.Next = d.Head
		d.Head.Pre = nn
		d.Head = nn
		d.length++
	}
}

func (d *Doubly[T]) InsertTail(data T) {
	nn := NewDoublyNode(data)

	switch {
	case d.isEmpty():
		d.Head = nn
		d.length++
		return

	case d.Head.Next == nil:
		d.Tail = nn
		d.Tail.Pre = d.Head
		d.Head.Next = d.Tail
		d.length++
		return

	default:
		d.Tail.Next = nn
		nn.Pre = d.Tail
		d.Tail = nn
		d.length++
	}
}

func (d *Doubly[T]) Insert(data T, position int) error {
	switch {
	case position < 0 || position > d.length:
		return ErrPositionOutOfRange

	case d.isEmpty():
		d.InsertHead(data)
		return nil

	case position == d.length:
		d.InsertTail(data)
		return nil

	default:
		cur := d.Head
		for i := 0; cur != nil && i < position; i++ {
			cur = cur.Next
		}
		nn := NewDoublyNode(data)
		nn.Pre = cur.Pre
		nn.Next = cur
		cur.Pre.Next = nn
		cur.Pre = nn
		d.length++
		return nil
	}
}

func (d *Doubly[T]) RemoveHead() error {
	switch {
	case d.isEmpty():
		return ErrListEmpty

	case d.Head.Next == nil:
		d.Head = nil
		d.length--
		return nil

	case d.Tail.Pre == d.Head && d.Head.Next == d.Tail:
		d.Head = d.Tail
		d.Head.Next = nil
		d.Tail = nil
		d.length--
		return nil

	default:
		d.Head = d.Head.Next
		d.length--
		return nil
	}
}

func (d *Doubly[T]) RemoveTail() error {
	switch {
	case d.isEmpty():
		return ErrListEmpty

	case d.Head.Next == nil:
		d.Head = nil
		d.length--
		return nil

	case d.Tail.Pre == d.Head && d.Head.Next == d.Tail:
		d.Tail = nil
		d.Head.Next = nil
		d.length--
		return nil

	default:
		d.Tail = d.Tail.Pre
		d.Tail.Next = nil
		d.length--
		return nil
	}
}

func (d *Doubly[T]) Remove(position int) error {
	switch {
	case position < 0 || position > d.length:
		return ErrPositionOutOfRange

	case d.isEmpty():
		return ErrListEmpty

	case position == 0:
		return d.RemoveHead()

	case position == d.length:
		return d.RemoveTail()

	default:
		cur := d.Head
		for i := 0; cur != nil && i < position; i++ {
			cur = cur.Next
		}
		cur.Pre.Next = cur.Next
		cur.Next.Pre = cur.Next.Pre
		d.length--
		return nil
	}
}

func (d *Doubly[T]) Transverse() {
	for cur := d.Head; cur != nil; cur = cur.Next {
		fmt.Printf("%v -> ", cur.Data)
	}
	fmt.Printf("nil")
	fmt.Println()
}

func (d *Doubly[T]) Len() int {
	return d.length
}

func (d *Doubly[T]) isEmpty() bool {
	return d.Head == nil || d.length == 0
}
