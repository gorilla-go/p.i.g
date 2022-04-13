package Bootstrap

import (
	"php-in-go/App/Exception"
	"php-in-go/Config"
	"php-in-go/Include/Contracts/Cache"
	"php-in-go/Include/Contracts/Debug"
	Http2 "php-in-go/Include/Contracts/Http"
	"php-in-go/Include/Contracts/Http/Log"
	"php-in-go/Include/Contracts/Http/Session"
	"php-in-go/Include/Contracts/Routing"
	Server2 "php-in-go/Include/Contracts/Server"
	Cache2 "php-in-go/Include/Foundation/Cache"
	Http4 "php-in-go/Include/Foundation/Http"
	Log2 "php-in-go/Include/Foundation/Http/Log"
	Session2 "php-in-go/Include/Foundation/Http/Session"
	"php-in-go/Include/Http"
	Routing2 "php-in-go/Include/Routing"
	"php-in-go/Routes"
)

type App struct {
	server           Server2.IServer
	router           Routing.IRouter
	exceptionHandler Debug.IExceptionHandler
	kernel           Http2.IKernel
	session          Session.ISession
	cache            Cache.ICache
	log              Log.ILog
}

func (a *App) Initializer(server Server2.IServer) {
	// set server.
	a.server = server

	// set http kernel route.
	a.router = &Routing2.Router{}
	a.router.Initializer(Routes.Route(), Config.Route())

	// set http exception handler.
	a.exceptionHandler = &Exception.Handler{}

	// cache.
	a.cache = &Cache2.MemoryCache{}
	a.cache.StartCacheManager()

	// set session driver.
	a.session = &Session2.Session{}
	a.session.StartSessionManager(a.cache)

	// set log server.
	a.log = &Log2.Log{}
	a.log.StartLogManager()

	// set http kernel.
	a.kernel = &Http4.Kernel{}
	a.kernel.Bootstrap(a)
}

func (a *App) Handle(request *Http.Request, response *Http.Response) {
	a.kernel.Handle(request, response)
}

func (a *App) GetServer() Server2.IServer {
	return a.server
}

func (a *App) GetRouter() Routing.IRouter {
	return a.router
}

func (a *App) GetLogger() Log.ILog {
	return a.log
}

func (a *App) GetSession() Session.ISession {
	return a.session
}

func (a *App) GetCache() Cache.ICache {
	return a.cache
}

func (a *App) GetExceptionHandler() Debug.IExceptionHandler {
	return a.exceptionHandler
}
