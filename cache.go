package main

import (
	"github.com/streamrail/concurrent-map"
)

var cache cmap.ConcurrentMap

func init() {
	cache = cmap.New()
}

func AddItemToCache(key string, value string) {
	cache.Set(key, value)
}

func GetItemFromCache(key string) string {
	if cache.Has(key) {
		if tmp, ok := cache.Get(key); ok {
			return tmp.(string)
		}
	}
	return ""
}

func removeItem(key string) {
	cache.Remove(key)
}
