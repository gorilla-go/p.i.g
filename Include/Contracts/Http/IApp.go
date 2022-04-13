package Http

import (
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	Server2 "php-in-go/Include/Contracts/Server"
	"php-in-go/Include/Http"
)

type IApp interface {
	Initializer(server Server2.IServer)
	Handle(request *Http.Request, response *Http.Response)
	GetServer() Server2.IServer
	GetExceptionHandler() Debug.IExceptionHandler
	GetRouter() Routing.IRouter
	GetLogger() Log.ILog
	GetCache() Cache.ICache
	GetSession() Session.ISession
}
