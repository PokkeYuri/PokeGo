package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       *sync.Mutex
	Counters map[string]cacheEntry
}

func NewCacheEntry(data []byte) cacheEntry {
	return cacheEntry{createdAt: time.Now(), val: data}
}
func NewCache(interval time.Duration) Cache {
	c := Cache{
		Counters: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go c.Reaploop(interval)
	return c
}
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.Counters[key] = cacheEntry{createdAt: time.Now(), val: value}
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.Counters[key]
	c.mu.Unlock()
	return val.val, ok
}
func (c *Cache) Reaploop(interval time.Duration) {
	for {
		nowTime := time.Now()
		c.mu.Lock()
		for k, v := range c.Counters {
			if v.createdAt.Add(interval).Before(nowTime) {
				delete(c.Counters, k)
			}
		}
		c.mu.Unlock()
	}
}
