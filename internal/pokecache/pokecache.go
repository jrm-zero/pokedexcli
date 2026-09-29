package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{}
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce, exists := c.cache[key]
	if exists {
		return ce.val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
			case <-ticker.C:
				c.mu.Lock()
				for key, ce := range c.cache {
					t1 := time.Now()
					if t1.Sub(ce.createdAt) > c.interval {
						delete(c.cache, key)
					}
				}
				c.mu.Unlock()
			default:
				continue
		}
	}
}