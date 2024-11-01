package linkedlist_test

import (
	"go-example/project/go-example/structures/linkedlist"
	"reflect"
	"testing"
)

type doublyLinkedListTest struct {
	data   int
	length int
}

func TestDoubly_InsertHead(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertHead(3)

		want := doublyLinkedListTest{data: 3, length: 1}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with one elements", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertHead(3)
		list.InsertHead(5)

		want := doublyLinkedListTest{data: 5, length: 2}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with multiple elements", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertHead(3)
		list.InsertHead(5)
		list.InsertHead(36)
		list.InsertHead(8)
		list.InsertHead(12)

		want := []int{12, 8, 36, 5, 3}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}

func TestDoubly_InsertTail(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertTail(3)

		want := doublyLinkedListTest{data: 3, length: 1}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with one elements", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertHead(3)
		list.InsertTail(5)

		want := doublyLinkedListTest{data: 5, length: 2}

		cur := list.Head
		for cur.Next != nil {
			cur = cur.Next
		}

		if cur.Data != want.data {
			t.Errorf("got: %v, want: %v", cur.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with multiple elements", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		list.InsertHead(3)
		list.InsertTail(5)
		list.InsertTail(36)
		list.InsertHead(8)
		list.InsertTail(12)

		want := []int{8, 3, 5, 36, 12}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestDoubly_Insert(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		want := doublyLinkedListTest{data: 3, length: 1}

		err := list.Insert(5, -1)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}

		err = list.Insert(3, 0)
		if err != nil {
			t.Error(err)
		}
		if list.Head.Data != want.data {

			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}
		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}

		err = list.Insert(4, 6)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}
	})

	t.Run("test with one elements", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()
		want := doublyLinkedListTest{data: 4, length: 1}

		err := list.Insert(5, -1)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}

		err = list.Insert(4, 0)
		if err != nil {
			t.Error(err)
		}
		if list.Head.Data != want.data {

			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}
		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}

		want = doublyLinkedListTest{data: 6, length: 2}

		err = list.Insert(6, 1)
		if err != nil {
			t.Error(err)
		}
		if list.Head.Next.Data != want.data {

			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}
		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}

		err = list.Insert(12, 4)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}
	})

	t.Run("test with multiple elements", func(t *testing.T) {

		list := linkedlist.NewDoubly[int]()

		err := list.Insert(5, -1)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}

		list.InsertHead(3)
		list.InsertTail(5)
		list.InsertTail(36)
		list.InsertHead(8)

		list.InsertTail(12)
		err = list.Insert(4, 2)
		if err != nil {
			t.Error(err)
		}

		want := []int{8, 3, 4, 5, 36, 12}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)

		}
		if list.Len() != 6 {
			t.Errorf("got: %v, want: %v", list.Len(), 6)
		}

		err = list.Insert(15, 8)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}
	})
}

func TestDoubly_RemoveHead(t *testing.T) {
	t.Run("test remove head", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()

		err := list.RemoveHead()
		if err != linkedlist.ErrListEmpty {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}
		if list.Len() != 0 {
			t.Errorf("got: %v, want: %v", list.Len(), 0)
		}

		list.InsertHead(3)
		list.InsertTail(5)
		list.InsertTail(36)
		list.InsertHead(8)

		err = list.RemoveHead()
		if err != nil {
			t.Error(err)
		}

		want := []int{3, 5, 36}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)

		}
		if list.Len() != 3 {
			t.Errorf("got: %v, want: %v", list.Len(), 3)
		}
	})
}

func TestDoubly_RemoveTail(t *testing.T) {
	t.Run("test remove tail", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()

		err := list.RemoveTail()
		if err != linkedlist.ErrListEmpty {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrListEmpty)
		}
		if list.Len() != 0 {
			t.Errorf("got: %v, want: %v", list.Len(), 0)
		}

		list.InsertTail(5)
		err = list.RemoveTail()
		if err != nil {
			t.Error(err)
		}

		list.InsertHead(3)
		list.InsertTail(5)
		list.InsertTail(36)
		list.InsertHead(8)

		err = list.RemoveTail()
		if err != nil {
			t.Error(err)
		}

		want := []int{8, 3, 5}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)

		}
		if list.Len() != 3 {
			t.Errorf("got: %v, want: %v", list.Len(), 3)
		}
	})
}

func TestDoubly_Remove(t *testing.T) {
	t.Run("test remove with position", func(t *testing.T) {
		list := linkedlist.NewDoubly[int]()

		err := list.Remove(-1)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}

		err = list.Remove(0)
		if err != linkedlist.ErrListEmpty {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrListEmpty)
		}

		list.InsertHead(3)
		list.InsertTail(5)
		list.InsertTail(36)
		list.InsertHead(8)
		list.InsertHead(4)
		list.InsertHead(12)
		list.InsertHead(96)

		err = list.Remove(3)
		if err != nil {
			t.Error(err)
		}

		want := []int{96, 12, 4, 3, 5, 36}
		got := []int{}
		cur := list.Head
		got = append(got, cur.Data)

		for cur.Next != nil {
			cur = cur.Next
			got = append(got, cur.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)

		}
		if list.Len() != 6 {
			t.Errorf("got: %v, want: %v", list.Len(), 6)
		}

		err = list.Remove(7)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}
	})
}
