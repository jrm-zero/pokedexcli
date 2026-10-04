package pokecache

import (
	"time"
	"sync"
	//"fmt"
)

type Cache struct {
	Cache map[string]CacheEntry
	interval time.Duration
	mu *sync.Mutex
}


type CacheEntry struct {
	CreatedAt time.Time
	Val []byte
}

func NewCache(interval time.Duration) Cache {
	m := make(map[string]CacheEntry)
	newCache := Cache{
		Cache: m,
		interval: interval,
		mu: &sync.Mutex{},
	}
	go newCache.reapLoop()
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce, exists := c.Cache[key]
	return ce.Val, exists
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	
	defer ticker.Stop()
	for range ticker.C {
				c.mu.Lock()
				for key, ce := range c.Cache {
					t1 := time.Now()
					if t1.Sub(ce.CreatedAt) > c.interval {
						delete(c.Cache, key)
					}
				}
				c.mu.Unlock()
	}
}