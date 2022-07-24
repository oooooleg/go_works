package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity  int
	queue     List
	items     map[Key]*ListItem
	lock      sync.Locker
	backItems map[interface{}]Key
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	item, ok := c.items[key]

	if ok {
		item.Value = value
		c.queue.MoveToFront(item)
		return true
	}

	c.items[key] = c.queue.PushFront(value)
	c.backItems[value] = key

	if c.queue.Len() > c.capacity {
		back := c.queue.Back()
		back.Prev.Next = nil
		backKey := c.backItems[back.Value]
		delete(c.items, backKey)
		delete(c.backItems, back.Value)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	item, ok := c.items[key]

	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(item)
	return item.Value, true
}

func (c *lruCache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.backItems = make(map[interface{}]Key)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity:  capacity,
		queue:     NewList(),
		items:     make(map[Key]*ListItem, capacity),
		lock:      &sync.Mutex{},
		backItems: make(map[interface{}]Key),
	}
}
