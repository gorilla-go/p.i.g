package Middleware

import (
	"net/http"
	"php-in-go/Config"
	"php-in-go/Include/Foundation/Util"
	"php-in-go/Include/Http"
	"php-in-go/Include/Routing"
	"time"
)

type SessionMiddleware struct {
}

func (m *SessionMiddleware) Handle(request *Http.Request, response *Http.Response, target *Routing.Target) bool {
	appConfig := Config.App()
	sessionKey := appConfig["sessionKey"].(string)

	cookie, err := request.Cookie(sessionKey)
	if err == http.ErrNoCookie || cookie.Value == "" {
		response.AddCookie(&http.Cookie{
			Name:    sessionKey,
			Value:   Util.Uuid(),
			Path:    "/",
			Expires: time.Now().Add(time.Second * time.Duration(appConfig["sessionExpire"].(int))),
		})
	}
	return true
}
