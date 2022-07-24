package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	count int
	head  *ListItem
	tail  *ListItem
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	added := &ListItem{Value: v, Next: l.head, Prev: nil}
	if l.count != 0 {
		l.head.Prev = added
	} else {
		l.tail = added
	}
	l.head = added
	l.count++
	return added
}

func (l *list) PushBack(v interface{}) *ListItem {
	added := &ListItem{Value: v, Next: nil, Prev: l.tail}

	if l.count != 0 {
		l.tail.Next = added
	} else {
		l.head = added
	}

	l.tail = added
	l.count++
	return added
}

func (l *list) Remove(i *ListItem) {
	l.count--

	if l.count == 0 {
		l.head, l.tail = nil, nil
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}

	if i == l.tail {
		l.tail = i.Prev
	}

	i.Prev.Next = i.Next
	l.head.Prev = i
	i.Prev, i.Next = nil, l.head
	l.head = i
}

func NewList() List {
	return new(list)
}
