package Cache

import "php-in-go/Include/Foundation/Cache"

type ICache interface {
	StartCacheManager()
	SetCache(key string, value interface{}, expire int, path string)
	GetCache(key string, path string) *Cache.Cache
	GetCachePath(path string) map[string]interface{}
	Clear(key string, path string)
	ClearPath(path string)
	ClearAll()
	CloseCacheManager()
}
