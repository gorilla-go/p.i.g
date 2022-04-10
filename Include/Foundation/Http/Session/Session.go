package Session

import (
	Config2 "php-in-go/Config"
	"php-in-go/Include/Contracts/Cache"
	"strings"
)

const SESSION_PREFIX = "session/"

type Session struct {
	cache Cache.ICache
}

func (s *Session) StartSessionManager(cache Cache.ICache) {
	s.cache = cache
}

func (s *Session) CloseSessionManager() {
}

func (s *Session) GetSession(str string) interface{} {
	return s.cache.GetCache(str, SESSION_PREFIX).Value
}

func (s *Session) SetSession(key string, v interface{}) {
	config := Config2.App()
	sessionExpire := config["sessionExpire"].(int)
	s.cache.SetCache(
		key,
		v,
		sessionExpire,
		SESSION_PREFIX,
	)
}

func (s *Session) GetSessionList() map[string]interface{} {
	c := make(map[string]interface{})
	group := s.cache.GetCachePath(SESSION_PREFIX)
	for k, v := range group {
		c[strings.Replace(k, SESSION_PREFIX, "", 1)] = v
	}
	return c
}
