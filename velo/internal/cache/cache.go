package cache

import (
	"sync"
	"time"

	"github.com/velo-api/velo/pkg/config"
)

type Cache struct {
	config config.CacheConfig
	items  map[string]*CacheItem
	mu     sync.RWMutex
}

type CacheItem struct {
	Value     []byte
	ExpiresAt time.Time
}

func New(cfg config.CacheConfig) *Cache {
	c := &Cache{
		config: cfg,
		items:  make(map[string]*CacheItem),
	}

	go c.cleanupLoop()

	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()

	if !ok || time.Now().After(item.ExpiresAt) {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

func (c *Cache) cleanupLoop() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, item := range c.items {
			if now.After(item.ExpiresAt) {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
