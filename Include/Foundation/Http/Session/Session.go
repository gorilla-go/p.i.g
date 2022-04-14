package Session

import (
	"php-in-go/Include/Contracts/Cache"
	"strings"
)

type Session struct {
	cache Cache.ICache
}

func (s *Session) StartSessionManager(cache Cache.ICache) {
	s.cache = cache
}

func (s *Session) CloseSessionManager() {
}

func (s *Session) GetSession(str string) interface{} {
	return s.cache.GetCache(str, "session/").Value
}

func (s *Session) SetSession(key string, v interface{}, expire int) {
	s.cache.SetCache(
		key,
		v,
		expire,
		"session/",
	)
}

func (s *Session) GetSessionList() map[string]interface{} {
	c := make(map[string]interface{})
	group := s.cache.GetCachePath("session/")
	for k, v := range group {
		c[strings.Replace(k, "session/", "", 1)] = v
	}
	return c
}
