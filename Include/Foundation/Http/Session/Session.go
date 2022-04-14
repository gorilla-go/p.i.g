package Session

import (
	"fmt"
	"net/http"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Http"
	"php-in-go/Include/Util"
	"time"
)

type Session struct {
	cache Cache.ICache
}

func (s *Session) StartSessionManager(cache Cache.ICache) {
	s.cache = cache
}

func (s *Session) CloseSessionManager() {
}

func (s *Session) GetSession(str string, request *Http.Request, response *Http.Response) interface{} {
	return s.cache.GetCache(
		str,
		fmt.Sprintf("session/%s/", s.getClientKey(request, response)),
	).Value
}

func (s *Session) SetSession(key string, v interface{}, request *Http.Request, response *Http.Response) {
	s.cache.SetCache(
		key,
		v,
		request.AppConfig["sessionExpire"].(int),
		fmt.Sprintf("session/%s/", s.getClientKey(request, response)),
	)
}

func (s *Session) Clear(request *Http.Request, response *Http.Response) {
	s.cache.ClearPath(fmt.Sprintf("session/%s/", s.getClientKey(request, response)))
}

func (s *Session) getClientKey(request *Http.Request, response *Http.Response) string {
	clientKey := request.AppConfig["sessionKey"].(string)
	cookie, err := request.Cookie(clientKey)
	cv := ""
	expire := request.AppConfig["sessionExpire"].(int)
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
