package cracking_the_coding_interview

import (
	"container/list"
	"fmt"
	"testing"
)

type LRUCache struct {
	cache map[int]*list.Element
	order *list.List
	cap   int
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*list.Element),
		order: list.New(),
		cap:   cap,
	}
}

func (c *LRUCache) Put(key int, value int) {
	if e, ok := c.cache[key]; ok {
		e.Value = value
		c.order.MoveToFront(e)
	} else {
		if c.order.Len() >= c.cap {
			c.removeOldest()
		}

		element := c.order.PushFront(key)
		c.cache[key] = element
	}
}

func (c *LRUCache) removeOldest() {
	oldest := c.order.Back()
	if oldest != nil {
		c.order.Remove(oldest)
		delete(c.cache, oldest.Value.(int))
	}
}

func (c *LRUCache) Get(key int) int {
	if e, ok := c.cache[key]; ok {
		c.order.MoveToFront(e)
		return e.Value.(int)
	}
	return -1
}

func Test_lru(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	fmt.Printf("cache=%v\n", cache)

	cache.Put(4, 4)
	fmt.Printf("cache=%v\n", cache.Get(1)) // -1 (已经被淘汰)
	fmt.Printf("cache=%v\n", cache.Get(2))
	fmt.Printf("cache=%v\n", cache.Get(3))
	fmt.Printf("cache=%v\n", cache.Get(4))
}
