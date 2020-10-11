package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Next  *listItem
	Prev  *listItem
}

type list struct {
	len   int
	front *listItem
	back  *listItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *listItem {
	return l.front
}

func (l *list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	listItem := listItem{v, nil, nil}
	l.len++
	if l.front != nil {
		l.front.Prev = &listItem
		listItem.Next = l.front
		l.front = &listItem
	} else {
		l.front = &listItem
		l.back = &listItem
	}
	return &listItem
}

func (l *list) PushBack(v interface{}) *listItem {
	listItem := listItem{v, nil, nil}
	l.len++
	if l.back != nil {
		l.back.Next = &listItem
		listItem.Prev = l.back
		l.back = &listItem
	} else {
		l.front = &listItem
		l.back = &listItem
	}

	return &listItem
}

func (l *list) Remove(i *listItem) {
	l.len--
	if l.Len() == 0 {
		i.Prev = nil
		i.Next = nil
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
		if l.back == i {
			l.back = i.Prev
		}
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
		if l.front == i {
			l.front = i.Next
		}
	}
}

func (l *list) MoveToFront(i *listItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if l.back != i {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	i.Next = l.front
	i.Prev = nil
	l.front = i
}

func NewList() List {
	return &list{}
}
