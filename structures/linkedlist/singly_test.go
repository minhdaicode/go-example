package linkedlist_test

import (
	"go-example/project/go-example/structures/linkedlist"
	"reflect"
	"testing"
)

type linkedlistTest struct {
	data   int
	length int
}

func TestSingly_InsertHead(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		list.InsertHead(3)

		want := linkedlistTest{
			3,
			1,
		}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with one elements", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		list.InsertHead(3)
		list.InsertHead(5)

		want := linkedlistTest{
			5,
			2,
		}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with multiple elements", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		list.InsertHead(3)
		list.InsertHead(5)
		list.InsertHead(36)
		list.InsertHead(8)
		list.InsertHead(12)

		want := linkedlistTest{
			12,
			5,
		}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})
}

func TestSingly_InsertTail(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		list.InsertTail(3)

		want := linkedlistTest{
			3,
			1,
		}

		if list.Head.Data != want.data {
			t.Errorf("got: %v, want: %v", list.Head.Data, want.data)
		}

		if list.Len() != want.length {
			t.Errorf("got: %v, want: %v", list.Len(), want.length)
		}
	})

	t.Run("test with one elements", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		list.InsertHead(3)
		list.InsertTail(5)

		want := linkedlistTest{
			5,
			2,
		}

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
		list := linkedlist.NewSingly[int]()
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

func TestSingly_Insert(t *testing.T) {
	t.Run("test with empty list", func(t *testing.T) {
		list := linkedlist.NewSingly[int]()
		want := linkedlistTest{
			3,
			1,
		}

		err := list.Insert(5, -1)
		if err != linkedlist.ErrPositionOutOfRange {
			t.Errorf("got error: %v, want error: %v", err, linkedlist.ErrPositionOutOfRange)
		}

		err = list.Insert(3, 0)
		if err != nil {
			t.Error(err)
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
		list := linkedlist.NewSingly[int]()
		list.InsertHead(3)
		list.Insert(5, 1)

		want := linkedlistTest{
			5,
			2,
		}

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
		list := linkedlist.NewSingly[int]()
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
