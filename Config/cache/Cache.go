package cache

func CacheConfig() map[string]interface{} {
	return map[string]interface{}{
		"defaultExpired": 60 * 60 * 24,
	}
}
