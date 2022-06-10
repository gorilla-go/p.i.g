package Session

import (
	"fmt"
	"net/http"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
	"php-in-go/Include/Util"
	"time"
)

type Session struct {
	cache  Cache.ICache
	config Config
}

func (s *Session) StartSessionManager(cache Cache.ICache, config Config) {
	s.cache = cache
	s.config = config
}

func (s *Session) CloseSessionManager() {
}

func (s *Session) GetSession(str string, request *Request.Request, response *Response.Response) interface{} {
	return s.cache.GetCache(
		str,
		fmt.Sprintf("session/%s/", s.SessionStart(request, response)),
	).Value
}

func (s *Session) SetSession(key string, v interface{}, request *Request.Request, response *Response.Response) {
	s.cache.SetCache(
		key,
		v,
		s.config.Expire,
		fmt.Sprintf("session/%s/", s.SessionStart(request, response)),
	)
}

func (s *Session) Clear(request *Request.Request, response *Response.Response) {
	s.cache.ClearPath(fmt.Sprintf("session/%s/", s.SessionStart(request, response)))
}

func (s *Session) SessionStart(request *Request.Request, response *Response.Response) string {
	clientKey := s.config.Name
	cookie, err := request.Cookie(clientKey)
	cv := ""
	expire := s.config.Expire
	if err == http.ErrNoCookie {
		uuid := Util.Uuid()
		response.AddCookie(&http.Cookie{
			Name:    clientKey,
			Value:   uuid,
			Path:    "/",
			Expires: time.Now().Add(time.Second * time.Duration(expire)),
		})
		cv = uuid
	} else {
		cv = cookie.Value
	}
	return cv
}
