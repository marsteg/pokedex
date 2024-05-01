package pokecache

import (
	"sync"
	"time"
)

/*I used a Cache struct to hold a map[string]cacheEntry and a
mutex to protect the map across goroutines.
A cacheEntry should be a struct with two fields:

createdAt - A time.Time that represents when the entry was created.
val - A []byte that represents the raw data we're caching.*/

type Cache struct {
	buffer map[string]cacheEntry
	mu     *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

//You'll probably want to expose a NewCache() function that creates
// a new cache with a configurable interval (time.Duration).

func NewCache(d time.Duration) Cache {
	NewCache := &Cache{
		buffer: make(map[string]cacheEntry),
		mu:     &sync.RWMutex{}, // Initialize the mutex here.
	}
	go NewCache.reapLoop(d)
	return *NewCache
}

// You'll want the following methods on the cache:
// ADD TO CACHE
// .Add() adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
func (c *Cache) Add(key string, val []byte) {
	var ce cacheEntry
	ce.createdAt = time.Now().UTC()
	ce.val = val
	c.mu.Lock()
	defer c.mu.Unlock()
	c.buffer[key] = ce
}

// GET FROM CACHE
// .Get() gets an entry from the cache. It should take a key (a string) and return a []byte and a bool.
// The bool should be true if the entry was found and false if it wasn't.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.buffer[key]
	return val.val, ok
}

//REAP LOOP
//The .reapLoop() method should be called when the cache is created (by the NewCache function).
// Each time an interval (the time.Duration passed to NewCache) passes it should remove any
// entries that are older than the interval. This makes sure that the cache doesn't
// grow too large over time. For example, if the interval is 5 seconds, and an entry was
// added 7 seconds ago, that entry should be removed.
//I used a time.Ticker to make this happen.

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.buffer {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.buffer, k)
		}
	}
}
