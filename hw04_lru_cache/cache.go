package hw04_lru_cache //nolint:golint,stylecheck
import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	purge() bool
	Clear()
}

type lruCache struct {
	lock     sync.Mutex
	capacity int
	queue    List
	items    map[Key]*listItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		item.Value.(*cacheItem).Value = value
		return true
	}

	item := l.queue.PushFront(&cacheItem{key, value})
	l.items[key] = item

	if l.queue.Len() > l.capacity {
		l.purge()
	}

	return false
}

func (l *lruCache) purge() bool {
	lastElem := l.queue.Back()

	if lastElem != nil {
		l.queue.Remove(lastElem)
		delete(l.items, lastElem.Value.(*cacheItem).Key)
		return true
	}

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	elem, ok := l.items[key]
	if ok {
		l.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).Value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.queue = NewList()
	l.items = make(map[Key]*listItem)
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		items:    make(map[Key]*listItem),
		capacity: capacity,
		queue:    NewList(),
	}
}
