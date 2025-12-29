// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cacheMap[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, false
}

func (c *Cache) reapLoop(ticker *time.Ticker) {
	for range ticker.C {
		for key, entry := range c.cacheMap {
			elapsed := time.Since(entry.createdAt)
			if elapsed > c.interval {
				delete(c.cacheMap, key)
			}
		}
	}
	defer ticker.Stop()
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{interval: interval}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	go c.reapLoop(ticker)

	return c
}
