package Cache

import (
	"context"
	"strings"
	"time"
)

type MemoryCache struct {
	pairContainer map[string]*Cache
	cancelLoop    context.CancelFunc
}

func (m *MemoryCache) StartCacheManager() {
	m.pairContainer = make(map[string]*Cache)
	cancel, cancelFunc := context.WithCancel(context.Background())
	m.cancelLoop = cancelFunc
	go func() {
		for true {
			select {
			case <-cancel.Done():
				return
			default:
				for name, cache := range m.pairContainer {
					if cache.Expire.Before(time.Now()) {
						delete(m.pairContainer, name)
					}
				}
				time.Sleep(time.Second * 60)
			}
		}
	}()
}

func (m *MemoryCache) SetCache(key string, value interface{}, expire int, path string) {
	m.pairContainer[path+key] = &Cache{
		Key:    key,
		Value:  value,
		Expire: time.Now().Add(time.Duration(expire) * time.Second),
	}
}

func (m *MemoryCache) Clear(key string, path string) {
	if _, exist := m.pairContainer[key+path]; exist {
		delete(m.pairContainer, key)
	}
}

func (m *MemoryCache) ClearPath(path string) {
	for k, _ := range m.pairContainer {
		if strings.HasPrefix(k, path) {
			delete(m.pairContainer, k)
		}
	}
}

func (m *MemoryCache) ClearAll() {
	m.pairContainer = make(map[string]*Cache)
}

func (m *MemoryCache) GetCache(key string, path string) *Cache {
	if v, exist := m.pairContainer[path+key]; exist && v.Expire.After(time.Now()) {
		return v
	}
	return nil
}

func (m *MemoryCache) GetCachePath(path string) map[string]interface{} {
	c := make(map[string]interface{})
	for k, v := range m.pairContainer {
		if strings.HasPrefix(k, path) {
			c[k] = v
		}
	}
	return c
}

func (m *MemoryCache) CloseCacheManager() {
	m.cancelLoop()
}
