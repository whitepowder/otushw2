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
	front *ListItem
	back  *ListItem
	size  int
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(value interface{}) *ListItem {
	item := &ListItem{value, l.front, nil}
	if l.front == nil && l.back == nil {
		l.front = item
	}
	if l.front != nil {
		l.front.Prev = item
		item.Next = l.front
		l.front = item
		item.Prev = nil
	}

	if l.size == 0 {
		l.back = item
	}
	l.size++
	return item
}

func (l *list) PushBack(value interface{}) *ListItem {
	item := &ListItem{value, nil, l.back}
	if l.back == nil {
		l.back = item
	}
	if l.back != nil {
		l.back.Next = item
		item.Prev = l.back
		l.back = item
		item.Next = nil
	}

	l.size++
	return item
}

func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.front = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.back = item.Prev
	}

	l.size--
	item.Prev, item.Next = nil, nil
}

func (l *list) MoveToFront(i *ListItem) {
	if l.size != 1 || i != l.front {
		l.Remove(i)
		l.PushFront(i.Value)
	}
}

func NewList() List {
	return new(list)
}
