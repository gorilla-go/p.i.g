package Config

import Cache2 "php-in-go/Include/Foundation/Cache"

func Cache() map[string]interface{} {
	return map[string]interface{}{
		"cacheDriver":    &Cache2.MemoryCache{},
		"defaultExpired": 60 * 60 * 24,
	}
}
