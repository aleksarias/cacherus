package main

import (
	"container/list"
	"errors"
	"log"
	"sync"
)

type Cache struct {
	size  int
	cache map[interface{}]*list.Element
	dList *list.List
	sync.Mutex
}

type Entry struct {
	key   interface{}
	value interface{}
}

func NewCache(size int) (*Cache, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive")
	}
	c := &Cache{
		size:  size,
		cache: make(map[interface{}]*list.Element),
		dList: list.New(),
	}
	return c, nil
}

func (c *Cache) Put(key interface{}, value interface{}) {

	c.Lock()
	defer c.Unlock()
	// Existing
	if ent, ok := c.cache[key]; ok {
		c.dList.MoveToFront(ent)
		ent.Value.(*Entry).value = value
	}

	// Add
	ent := &Entry{key, value}
	entry := c.dList.PushFront(ent)
	c.cache[key] = entry

	// Remove if necessary
	if c.dList.Len() > c.size {
		ent := c.dList.Back()
		if ent != nil {
			c.dList.Remove(ent)
			kv := ent.Value.(*Entry)
			delete(c.cache, kv.key)
			log.Printf("evicted %v: %v", kv.key, kv.value)
		}
	}
}

func (c *Cache) Get(key interface{}) (interface{}, error) {
	c.Lock()
	defer c.Unlock()
	ent, ok := c.cache[key]
	if ok {
		c.dList.MoveToFront(ent)
		return ent.Value.(*Entry).value, nil
	}
	return nil, errors.New("key not found")
}
