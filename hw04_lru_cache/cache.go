package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue *cacheItem
	items *cacheItem
}

func (l lruCache) Set(key string, value interface{}) bool {
	panic("implement me")
}

func (l lruCache) Get(key string) (interface{}, bool) {
	panic("implement me")
}

func (l lruCache) Clear() {
	panic("implement me")
}

type cacheItem struct {
}

func NewCache(capacity int) Cache {
	return &lruCache{}
}
