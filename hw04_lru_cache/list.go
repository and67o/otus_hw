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

func (l list) Len() int {
	return l.len
}

func (l list) Front() *listItem {
	return l.front
}

func (l list) Back() *listItem {
	return l.back
}

func (l list) PushFront(v interface{}) *listItem {
	listItem := listItem{v, nil, l.Back()}
	l.len++
	if l.Front() != nil {
		l.front.Prev = &listItem
		listItem.Next = l.back
		l.back = &listItem
	} else {
		l.front = &listItem
		l.back = &listItem
	}
	return &listItem
}

func (l list) PushBack(v interface{}) *listItem {
	listItem := listItem{v, nil, l.Back()}
	if l.Back() != nil {
		l.back.Next = &listItem
		listItem.Prev = l.back
		l.front = &listItem
	} else {
		l.front = &listItem
		l.back = &listItem
	}
	l.len++
	return &listItem
}

func (l list) Remove(i *listItem) {
	panic("implement me")
}

func (l list) MoveToFront(i *listItem) {
	panic("implement me")
}

func NewList() List {
	return &list{}
}
