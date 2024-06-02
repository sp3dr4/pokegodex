package cache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Printf("cache-add %s\n", key)
	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Lock()
	c.entries[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	if !ok {
		fmt.Printf("cache-miss %s\n", key)
		return nil, false
	}
	fmt.Printf("cache-hit %s\n", key)
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for t := range ticker.C {
		fmt.Printf("reapLoop %v\n", t)
		c.mu.Lock()
		for key, entry := range c.entries {
			if t.Sub(entry.createdAt) > interval {
				fmt.Printf("\tdeleting key %s\n", key)
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: map[string]cacheEntry{},
		mu:      sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return &c
}
