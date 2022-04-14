package Session

import (
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Http"
)

type ISession interface {
	StartSessionManager(cache Cache.ICache)
	CloseSessionManager()
	GetSession(s string, request *Http.Request, response *Http.Response) interface{}
	SetSession(key string, v interface{}, request *Http.Request, response *Http.Response)
	Clear(request *Http.Request, response *Http.Response)
}
