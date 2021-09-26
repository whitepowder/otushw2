package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mtex     sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mtex.Lock()
	defer c.mtex.Unlock()
	element, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(element)
		element.Value = cacheItem{key, value}
		return true
	}

	if c.queue.Len() == c.capacity {
		if element := c.queue.Back(); element != nil {
			c.queue.Remove(element)
		}
	}

	element = c.queue.PushFront(cacheItem{key, value})
	c.items[key] = element

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mtex.Lock()
	defer c.mtex.Unlock()
	element, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(element)
		return element.Value.(cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mtex.Lock()
	c.items = make(map[Key]*ListItem)
	c.queue = NewList()
	c.mtex.Unlock()
}
