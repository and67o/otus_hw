package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("empty logic", func(t *testing.T) {
		c := NewCache(3)
		wasInCache := c.Set("111", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("222", 200)
		require.False(t, wasInCache)

		c.Clear()

		val, ok := c.Get("111")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("used logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("1", 100)
		c.Set("2", 200)
		c.Set("3", 300)

		_, ok := c.Get("2")
		require.True(t, ok)

		c.Set("3", 301)
		c.Set("3", 302)
		_, ok1 := c.Get("3")
		require.True(t, ok1)

		c.Set("2", 101)
		c.Set("4", 400)

		keys := make([]Key, 0, 3)
		for key := range c.List() {
			keys = append(keys, key)
		}

		require.Subset(t, []Key{"2", "3", "4"}, keys)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		wasInCache := c.Set("111", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("222", 200)
		require.False(t, wasInCache)

		wasInCache = c.Set("333", 300)
		require.False(t, wasInCache)

		wasInCache = c.Set("444", 400)
		require.False(t, wasInCache)

		val, ok := c.Get("111")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("knockout logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("111", 100)
		c.Set("222", 222)
		c.Set("333", 333)
		c.Set("444", 444)

		val, ok := c.Get("111")
		require.False(t, ok)
		require.Nil(t, val)

		keys := make([]Key, 0, 3)
		for key := range c.List() {
			keys = append(keys, key)
		}
		require.Subset(t, []Key{"222", "333", "444"}, keys)
	})
}

func TestCacheMultithreading(t *testing.T) {

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
