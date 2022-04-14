package Middleware

import (
	"net/http"
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
	"php-in-go/Include/Util"
	"time"
)

type SessionMiddleware struct {
}

func (m *SessionMiddleware) Handle(request *Http.Request, response *Http.Response, target *Routing.Target) bool {
	sessionKey := request.AppConfig["sessionKey"].(string)

	cookie, err := request.Cookie(sessionKey)
	if err == http.ErrNoCookie || cookie.Value == "" {
		response.AddCookie(&http.Cookie{
			Name:    sessionKey,
			Value:   Util.Uuid(),
			Path:    "/",
			Expires: time.Now().Add(time.Second * time.Duration(request.AppConfig["sessionExpire"].(int))),
		})
	}
	return true
}
