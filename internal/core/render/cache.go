package render

import "sync"

// imageCache is a simple in memory cache keyed by image URL
// Entries persist for the lifetime of the session
type imageCache struct {
	mu    sync.RWMutex
	store map[string]string
}

func newImageCache() *imageCache {
	return &imageCache{store: make(map[string]string)}
}

func (c *imageCache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.store[key]
	return v, ok
}

func (c *imageCache) set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

// globalCache is the package level cache instance used by ImageToANSI
var globalCache = newImageCache()
