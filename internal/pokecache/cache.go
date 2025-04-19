package pokecache

import (
	"sync"
	"time"
)

var Blue = "\033[34m"
var Reset = "\033[0m"

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	entries map[string]cacheEntry
	sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, b := c.entries[key]
	if !b {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	// fmt.Println(Blue + "Ticker Started" + Reset)
	for range ticker.C {
		c.Lock()
		for key, entry := range c.entries {
			curTime := time.Now()
			lapsedTime := curTime.Sub(entry.createdAt)
			if lapsedTime > interval {
				// Remove the entry
				delete(c.entries, key)
				continue
			}

		}
		c.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
	}
	// fmt.Println(Blue + "Cache Created" + Reset)
	go c.reapLoop(interval)
	return &c
}
