package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("steps", func(t *testing.T) {
		l := NewList()

		l.PushFront(22)
		require.Equal(t, 1, l.Len())
		require.Equal(t, 22, l.Front().Value)
		require.Equal(t, l.Front(), l.Back())
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Front().Prev)

		l.Remove(l.Front())
		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())

		l.PushBack(33)
		pushedLast := l.PushBack(34)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 33, l.Front().Value)
		require.Equal(t, 34, l.Back().Value)
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Equal(t, l.Back(), l.Front().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)

		l.MoveToFront(pushedLast)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 34, l.Front().Value)
		require.Equal(t, 33, l.Back().Value)
		require.Equal(t, l.Front(), l.Back().Prev)
		require.Equal(t, l.Back(), l.Front().Next)
		require.Nil(t, l.Front().Prev)
		require.Nil(t, l.Back().Next)

		l.PushFront(77)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 77, l.Front().Value)
		require.Equal(t, 33, l.Back().Value)
		require.Equal(t, l.Front().Next, l.Back().Prev)
		require.Equal(t, 34, l.Front().Next.Value)

		l.MoveToFront(l.Front().Next)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 34, l.Front().Value)
		require.Equal(t, 77, l.Front().Next.Value)
		require.Equal(t, 33, l.Front().Next.Next.Value)
		require.Equal(t, l.Front().Next.Next.Value, l.Back().Value)
	})
}
