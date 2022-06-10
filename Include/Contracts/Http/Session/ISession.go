package Session

import (
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http/Request"
	"php-in-go/Include/Http/Response"
)

type ISession interface {
	StartSessionManager(cache Cache.ICache, config Session.Config)
	CloseSessionManager()
	GetSession(s string, request *Request.Request, response *Response.Response) interface{}
	SetSession(key string, v interface{}, request *Request.Request, response *Response.Response)
	Clear(request *Request.Request, response *Response.Response)
	SessionStart(request *Request.Request, response *Response.Response) string
}
