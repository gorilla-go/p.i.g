package Session

import (
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http"
)

type ISession interface {
	StartSessionManager(cache Cache.ICache, config Session.Config)
	CloseSessionManager()
	GetSession(s string, request *Http.Request, response *Http.Response) interface{}
	SetSession(key string, v interface{}, request *Http.Request, response *Http.Response)
	Clear(request *Http.Request, response *Http.Response)
}
