package App

import (
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	"php-in-go/Include/Http"
)

type IApp interface {
	Initializer()
	Close()
	Handle(request *Http.Request, response *Http.Response)
	GetExceptionHandler() Debug.IExceptionHandler
	GetRouter() Routing.IRouter
	GetLogger() Log.ILog
	GetCache() Cache.ICache
	GetSession() Session.ISession
	GetConfigs() map[string]map[string]interface{}
}
