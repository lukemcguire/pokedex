package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}
	// background reaping process
	go cache.reapLoop()

	return cache
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[url]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Add(url string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	c.entries[url] = entry

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now().UTC()
		cutoff := now.Add(-c.interval)
		c.reap(cutoff)
	}
}

func (c *Cache) reap(cutoff time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}
